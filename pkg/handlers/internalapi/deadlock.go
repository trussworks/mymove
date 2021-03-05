package internalapi

import (
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"go.uber.org/zap"
	"time"

	deadlockop "github.com/transcom/mymove/pkg/gen/internalapi/internaloperations/dead"
	"github.com/transcom/mymove/pkg/handlers"
)

// FetchAccessCodeHandler fetches an access code associated with a service member
type ForceDeadlockHandler struct {
	handlers.HandlerContext
}

// Handle fetches the access code for a service member
func (h ForceDeadlockHandler) Handle(params deadlockop.ForceDeadlockParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	logger := h.LoggerFromContext(ctx)
	fmt.Println("I'm in the deadlock handler🥠🥠🥠🥠🥠🥠🥠")
	c := make(chan bool)
	logger.Info("Before deadlock", zap.Time("time: ", time.Now()))
	<-c
	fmt.Println("vim-go past deadlock!")
	return deadlockop.NewForceDeadlockOK()
}
