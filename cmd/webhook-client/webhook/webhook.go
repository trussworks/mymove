package webhook

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gofrs/uuid"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/transcom/mymove/cmd/webhook-client/utils"
	"github.com/transcom/mymove/pkg/models"
)

//webhookMessage is the body of our request
type webhookMessage struct {
	// Message sent
	// Required: true
	Message string `json:"message"`
	// Notification ID
	ID uuid.UUID `json:"id"`
}

// Engine encapsulates the services used by the webhook notification engine
type Engine struct {
	Connection      *pop.Connection
	Logger          utils.Logger
	Client          utils.WebhookClientPoster
	Cmd             *cobra.Command
	PeriodInSeconds int
}

func (eng *Engine) processNotifications(notifications []models.WebhookNotification, subscriptions []models.WebhookSubscription) {
	for _, notif := range notifications {

		// search for subscription
		foundSub := false
		for _, sub := range subscriptions {
			if sub.EventKey == notif.EventKey {
				foundSub = true
				// If found, send  to subscription
				err := eng.sendOneNotification(&notif, &sub)
				if err != nil {
					eng.Logger.Error("Webhook Notification send failed", zap.Error(err))
					return
				}
				continue
			}
		}
		if foundSub == false {
			// MYTODO: Update notification status to skipped, once that's available. Currently updating to pending.
			eng.Logger.Debug("No subscription found for notification event.", zap.String("eventKey", notif.EventKey))
			notif.Status = models.WebhookNotificationPending
			err := eng.updateNotification(&notif)
			if err != nil {
				eng.Logger.Error("Notification update failed", zap.Error(err))
			}
		}

	}
}

func (eng *Engine) updateNotification(notif *models.WebhookNotification) error {
	// Update notification (status is updated by sendOneNotification)
	verrs, err := eng.Connection.ValidateAndUpdate(notif)
	if verrs != nil && verrs.HasAny() {
		err = errors.New(verrs.Error())
		eng.Logger.Error(err.Error())
		return err
	}
	if err != nil {
		eng.Logger.Error(err.Error())
		return err
	}
	return nil
}

// sendOneNotification sends the notification and marks the notification as sent or failed
func (eng *Engine) sendOneNotification(notif *models.WebhookNotification, sub *models.WebhookSubscription) error {
	logger := eng.Logger

	// Construct notification to send
	message := fmt.Sprintf("There was a %s notification, id: %s, moveTaskOrderID: %s",
		string(notif.EventKey), notif.ID.String(), notif.MoveTaskOrderID.String())
	payload := &webhookMessage{
		Message: message,
		ID:      notif.ID,
	}
	json, err := json.Marshal(payload)
	if err != nil {
		notif.Status = models.WebhookNotificationFailed
		eng.updateNotification(notif)
		logger.Error("Error creating payload:", zap.Error(err))
		return err
	}

	// Post the notification
	url := sub.CallbackURL
	resp, body, err := eng.Client.Post(json, url)

	// Check for error making the request
	if err != nil {
		notif.Status = models.WebhookNotificationFailed
		eng.updateNotification(notif)
		logger.Error("Failed to send, error making request:", zap.Error(err))
		return err
	}
	// Check for error response from server
	if resp.StatusCode != 200 {
		notif.Status = models.WebhookNotificationFailed
		eng.updateNotification(notif)
		errmsg := fmt.Sprintf("Failed to send. Response Status: %s. Body: %s", resp.Status, string(body))
		err = errors.New(errmsg)
		logger.Error("db-webhook-notify: Failed to send notification", zap.Error(err))
		return err
	}

	notif.Status = models.WebhookNotificationSent
	eng.updateNotification(notif)
	logger.Info("Notification successfully sent:", zap.String("Status", resp.Status), zap.String("Body", string(body)))
	return nil

}

// Run runs the engine once. This happens periodically and is called by Start()
// It collects all the pending notifications and active subscriptions in the db
// and starts processing them.
// If a new notification or subscription were to be adding during the course of one run
// by the Milmove server, it would only be processed on the next call of run().
func (eng *Engine) run() error {

	logger := eng.Logger
	// Read all notifications
	notifications := []models.WebhookNotification{}
	err := eng.Connection.Order("created_at asc").Where("status = ?", models.WebhookNotificationPending).All(&notifications)

	if err != nil {
		logger.Error("Error:", zap.Error(err))
		return err
	}
	logger.Debug("Success!", zap.Int("Num notifications found", len(notifications)))

	// If none, return
	if len(notifications) == 0 {
		return nil
	}

	// If there are notifications, get subscriptions
	subscriptions := []models.WebhookSubscription{}
	err = eng.Connection.Where("status = ?", models.WebhookSubscriptionStatusActive).All(&subscriptions)

	if err != nil {
		logger.Error("Error:", zap.Error(err))
		return err
	}
	logger.Debug("Success!", zap.Int("Num subscriptions found", len(subscriptions)))

	// If none, return
	if len(notifications) == 0 {
		return nil
	}

	// MYTODO: Maybe want to reorganize subs in memory for faster access
	// process notifications
	eng.processNotifications(notifications, subscriptions)
	return nil
}

// Start starts the timer for the webhook engine
// The process will run once every period to send pending notifications
// The period is defined in the WebhookEngine PeriodInSeconds
func (eng *Engine) Start() error {

	// Set timer tick
	t := time.Tick(time.Duration(eng.PeriodInSeconds) * time.Second)

	// Run once prior to first wait period
	eng.run()
	// Run on each timer tick
	for range t {
		eng.run()
	}

	// todo when and how should we kick out of this function?
	return nil
}
