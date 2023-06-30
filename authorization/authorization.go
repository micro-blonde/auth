package authorization

import (
	"context"
	"sync"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/micro-blonde/auth/account"
)

type Authorization[T account.Model] interface {
	gateway.Authorization

	WithSession(s *Session[T]) Authorization[T]
	GetSession() *Session[T]
	UpdateSession(ctx context.Context, session *Session[T]) errors.Error

	GetAccount() *account.Account[T]

	SetCaptchaVerified(verified bool)
	IsCaptchaVerified() bool
}

type authorization[T account.Model] struct {
	valuesMtx *sync.RWMutex
	values    map[string]any

	authenticator *authenticator[T]

	session *Session[T]

	captchaVerified bool
}

func newAuthorization[T account.Model](request gateway.Request,
	authenticator *authenticator[T], session *Session[T]) (auth Authorization[T]) {
	auth0 := request.GetAuthorization()
	if auth0 != nil {
		auth = auth0.(Authorization[T])
		return auth.WithSession(session)
	}
	auth = &authorization[T]{
		valuesMtx:     new(sync.RWMutex),
		values:        make(map[string]any),
		authenticator: authenticator,
		session:       session,
	}
	return auth
}

func (a *authorization[T]) Set(key string, value any) {
	a.valuesMtx.Lock()
	defer a.valuesMtx.Unlock()

	a.values[key] = value
}

func (a *authorization[T]) Get(key string) any {
	a.valuesMtx.RLock()
	defer a.valuesMtx.RUnlock()

	return a.values[key]
}

func (a *authorization[T]) GetApplicantId() any {
	if a.session == nil {
		return nil
	}
	return a.session.Account.Id
}

func (a *authorization[T]) GetApplicant() any {
	if a.session == nil {
		return nil
	}
	return a.session.Account
}

func (a *authorization[T]) GetAccount() *account.Account[T] {
	if a.session == nil {
		return nil
	}
	return &a.session.Account
}

func (a *authorization[T]) SetCaptchaVerified(verified bool) {
	a.captchaVerified = verified
}

func (a *authorization[T]) IsCaptchaVerified() bool {
	return a != nil && a.captchaVerified
}
