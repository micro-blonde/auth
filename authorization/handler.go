package authorization

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
)

type handler struct {
	gateway.Responder
	handleFunc gateway.HandleFunc
}

func (h *authenticator[T]) newHandler(responder gateway.Responder,
	handleFunc gateway.HandleFunc) gateway.Handler {
	r := &handler{
		Responder:  responder,
		handleFunc: handleFunc,
	}
	return r
}

func (h *handler) Handle(request gateway.Request) (any, errors.Error) {
	return h.handleFunc(request)
}
