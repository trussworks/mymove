package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/transcom/mymove/cmd/webhook-client/utils"
	"github.com/transcom/mymove/cmd/webhook-client/webhook"

	"go.uber.org/zap"
)

const (
	// PeriodFlag indicates how often to check the db in seconds
	PeriodFlag string = "period"
	// MaxRetriesFlag indicates how many times to immediately retry
	MaxRetriesFlag string = "max-retries"
)

// Init flags specific to this command
func initWebhookNotifyFlags(flag *pflag.FlagSet) {
	flag.Int(PeriodFlag, 5, "Period in secs to check for notifications")
	flag.Int(MaxRetriesFlag, 3, "Number of times to immediately retry")
	flag.SortFlags = false
}

func webhookNotify(cmd *cobra.Command, args []string) error {
	v := viper.New()

	// Parse flags from command line
	err := utils.ParseFlags(cmd, v, args)
	if err != nil {
		return err
	}

	// Validate all arguments passed in including DB, CAC, etc...
	// Also this opens the db connection and creates a logger
	db, logger, err := InitRootConfig(v)
	if err != nil {
		logger.Error("invalid configuration", zap.Error(err))
		return err
	}

	// Create the client and open the cacStore (which has to stay open?)
	runtime, cacStore, errCreateClient := utils.CreateClient(v)

	if errCreateClient != nil {
		logger.Error("Error creating client:", zap.Error(errCreateClient))
		return errCreateClient
	}

	// Defer closing the store until after the api call has completed
	if cacStore != nil {
		defer cacStore.Close()
	}

	// Create a webhook engine
	webhookEngine := webhook.Engine{
		DB:                  db,
		Logger:              logger,
		Client:              runtime,
		PeriodInSeconds:     v.GetInt(PeriodFlag),
		MaxImmediateRetries: v.GetInt(MaxRetriesFlag),
		SeverityThresholds:  []int{60},
	}

	// Start polling the db for changes
	go webhookEngine.Start()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")
	if err = db.Close(); err == nil {
		logger.Info("Db connection closed")
	} else {
		logger.Error("Db connection close failed", zap.Error(err))
	}

	log.Println("Listener exiting")
	return nil
}