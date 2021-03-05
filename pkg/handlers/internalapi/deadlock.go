package internalapi

import (
	"fmt"

	"github.com/go-openapi/runtime/middleware"

	deadlockop "github.com/transcom/mymove/pkg/gen/internalapi/internaloperations/dead"
	"github.com/transcom/mymove/pkg/handlers"
)

// FetchAccessCodeHandler fetches an access code associated with a service member
type ForceDeadlockHandler struct {
	handlers.HandlerContext
}

// Handle fetches the access code for a service member
func (h ForceDeadlockHandler) Handle(params deadlockop.ForceDeadlockParams) middleware.Responder {
	fmt.Println("I'm in the deadlock handler🥠🥠🥠🥠🥠🥠🥠")
	fmt.Println("entering deadlock")
	c := make(chan bool)
	<-c
	fmt.Println("vim-go past deadlock!")
	return deadlockop.NewForceDeadlockOK()
}
