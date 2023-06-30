package authorization

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
)

func (h *authenticator[T]) Authenticate(request gateway.Request,
	token string) (Authorization[T], errors.Error) {
	key := h.config.AccessKeyPrefix + token

	session := new(Session[T])
	err := h.cache.Load(request.GetContext(), key, session)
	if err != nil {
		if err.IsType(errors.TypeNotFound) {
			return nil, errors.Unauthorized(err)
		}
		return nil, err
	}
	session.Key = key

	auth := newAuthorization(request, h, session)
	request.SetAuthorization(auth)
	return auth, nil
}

func (h *authenticator[T]) authenticate(
	request gateway.Request) (Authorization[T], errors.Error) {
	header := request.GetHeader("authorization")
	if len(header) < 7 {
		return nil, errors.Unauthorized()
	}

	return h.Authenticate(request, header[7:])
}
