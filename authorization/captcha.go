package authorization

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
)

func (h *authenticator[T]) verifyCaptcha(request gateway.Request) errors.Error {
	if h.config.Captcha.Debug {
		return nil
	}

	secret := request.GetHeader("X-Captcha-Secret")
	if secret == "" {
		return InvalidCaptchaError
	}

	passcode := request.GetHeader("X-Captcha-PassCode")
	if passcode == "" {
		return EmptyCaptchaPasscodeError
	}

	captchaKey := "captcha_" + secret
	captcha := make(map[string]any)
	err := h.cache.Load(request.GetContext(), captchaKey, &captcha)
	if err != nil {
		return InvalidCaptchaPasscodeError
	}

	err = h.cache.Delete(request.GetContext(), captchaKey)
	if err != nil {
		return InvalidCaptchaPasscodeError
	}

	return InvalidCaptchaPasscodeError
}

func (h *authenticator[T]) setCaptchaVerified(request gateway.Request, verified bool) {
	authModel := request.GetAuthorization()
	if authModel == nil {
		authModel = newRequestAuthorization(request, h, nil)
		request.SetAuthorization(authModel)
	}
	authModel.(Authorization[T]).SetCaptchaVerified(verified)
}
