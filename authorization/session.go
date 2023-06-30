package authorization

import (
	"context"
	"time"

	"github.com/ginger-core/errors"
	"github.com/micro-blonde/auth/account"
)

type Session[T account.Model] struct {
	Key       string `json:"-"`
	Id        string
	CreatedAt time.Time

	AccessToken    string
	AccessTokenExp time.Duration

	RefreshToken    string
	RefreshTokenExp time.Duration

	Account account.Account[T]

	Scopes []string

	Meta Meta
}

func (a *authorization[T]) WithSession(s *Session[T]) Authorization[T] {
	a.session = s
	return a
}

func (a *authorization[T]) GetSession() *Session[T] {
	return a.session
}

func (a *authorization[T]) UpdateSession(
	ctx context.Context, session *Session[T]) errors.Error {
	exp, err := a.authenticator.cache.GetExpiration(ctx, session.Key)
	if err != nil {
		return errors.New(err).
			WithTrace("UpdateSession.GetExpiration.err")
	}
	if exp <= 0 {
		return errors.Internal().
			WithTrace("UpdateSession.GetExpiration.exp0")
	}
	return a.authenticator.cache.MarshalStore(ctx, session.Key, session, exp)
}

func (s *Session[T]) GetMeta() Meta {
	if s.Meta == nil {
		s.Meta = make(Meta)
	}
	return s.Meta
}

func (s *Session[T]) HasScope(scopes ...string) bool {
	for _, s1 := range scopes {
		for _, s0 := range s.Scopes {
			if s1 == s0 {
				return true
			}
		}
	}
	return false
}
