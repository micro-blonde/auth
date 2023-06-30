package authorization

import (
	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/repository"
	"github.com/ginger-repository/redis"
	"github.com/micro-blonde/auth/account"
)

type Authenticator[T account.Model] interface {
	ShouldVerifyCaptcha() gateway.Handler
	MustVerifyCaptcha() gateway.Handler

	Authenticate(request gateway.Request,
		token string) (Authorization[T], errors.Error)

	ShouldAuthenticate() gateway.Handler
	MustAuthenticate() gateway.Handler
	MustHaveScope(scopes ...string) gateway.Handler
	MustNotHaveScope(scopes ...string) gateway.Handler
}

type authenticator[T account.Model] struct {
	server gateway.Server
	config config
	cache  repository.Cache
}

func New[T account.Model](server gateway.Server,
	registry registry.Registry) Authenticator[T] {
	h := &authenticator[T]{
		server: server,
	}
	if err := registry.Unmarshal(&h.config); err != nil {
		panic(err)
	}
	h.config.initialize()

	db := redis.NewRepository(registry.ValueOf("redis"))
	if err := db.Initialize(); err != nil {
		panic(err)
	}

	h.cache = redis.NewCache(db)
	return h
}

func (h *authenticator[T]) ShouldVerifyCaptcha() gateway.Handler {
	return h.newHandler(h.server.GetController(), h.shouldVerifyCaptchaHandle)
}

func (h *authenticator[T]) MustVerifyCaptcha() gateway.Handler {
	return h.newHandler(h.server.GetController(), h.mustVerifyCaptchaHandle)
}

func (h *authenticator[T]) ShouldAuthenticate() gateway.Handler {
	return h.newHandler(h.server.GetController(), h.shouldAuthenticateHandle)
}

func (h *authenticator[T]) MustAuthenticate() gateway.Handler {
	return h.newHandler(h.server.GetController(), h.mustAuthenticateHandle)
}

func (h *authenticator[T]) MustHaveScope(scopes ...string) gateway.Handler {
	if len(scopes) == 0 {
		return h.MustAuthenticate()
	}
	return h.newHandler(h.server.GetController(),
		h.mustHaveScopeHandle(scopes...))
}

func (h *authenticator[T]) MustNotHaveScope(scopes ...string) gateway.Handler {
	if len(scopes) == 0 {
		return h.MustAuthenticate()
	}
	return h.newHandler(h.server.GetController(),
		h.mustNotHaveScopeHandle(scopes...))
}
