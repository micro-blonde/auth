package authorization

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
)

func (h *authenticator[T]) shouldVerifyCaptchaHandle(
	request gateway.Request) (any, errors.Error) {
	err := h.verifyCaptcha(request)
	if err != nil {
		h.setCaptchaVerified(request, false)
		return nil, nil
	}

	h.setCaptchaVerified(request, true)
	return nil, nil
}

func (h *authenticator[T]) mustVerifyCaptchaHandle(
	request gateway.Request) (any, errors.Error) {
	err := h.verifyCaptcha(request)
	if err != nil {
		return nil, err
	}

	h.setCaptchaVerified(request, true)
	return nil, nil
}

func (h *authenticator[T]) shouldAuthenticateHandle(
	request gateway.Request) (any, errors.Error) {
	_, _ = h.authenticate(request)
	return nil, nil
}

func (h *authenticator[T]) mustAuthenticateHandle(
	request gateway.Request) (any, errors.Error) {
	_, err := h.authenticate(request)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (h *authenticator[T]) mustHaveScopeHandle(
	scopes ...string) gateway.HandleFunc {
	return func(request gateway.Request) (any, errors.Error) {
		auth, err := h.authenticate(request)
		if err != nil {
			return nil, err
		}

		sess := auth.GetSession()
		found := false
		for _, s := range scopes {
			for _, s2 := range sess.Scopes {
				if s == s2 {
					found = true
					break
				}
			}
			if found {
				break
			}
		}
		if !found {
			return nil, errors.Forbidden().
				WithContext(request.GetContext()).
				WithDetail(errors.NewDetail().
					With("scopes", scopes)).
				WithDesc("scope not found")
		}
		return nil, nil
	}
}

func (h *authenticator[T]) mustNotHaveScopeHandle(
	scopes ...string) gateway.HandleFunc {
	return func(request gateway.Request) (any, errors.Error) {
		auth, err := h.authenticate(request)
		if err != nil {
			return nil, err
		}

		sess := auth.GetSession()
		for _, s := range scopes {
			found := false
			for _, s2 := range sess.Scopes {
				if s == s2 {
					found = true
					break
				}
			}
			if found {
				return nil, errors.Forbidden().
					WithContext(request.GetContext()).
					WithDesc("scope `" + s + "` found")
			}
		}
		return nil, nil
	}
}
