package authorization

import "github.com/ginger-core/errors"

var (
	InvalidCaptchaError = errors.Validation().
				WithId("InvalidCaptcha").
				WithCode(CaptchaErrorCode).
				WithMessage("Error getting captcha information, please fill captcha information")
	EmptyCaptchaPasscodeError = errors.Validation().
					WithId("EmptyCaptchaPasscode").
					WithCode(CaptchaErrorCode).
					WithMessage("Error getting captcha passcode, please fill captcha information")
	InvalidCaptchaPasscodeError = errors.Validation().
					WithId("InvalidCaptchaPasscode").
					WithCode(CaptchaErrorCode).
					WithMessage("Error verifying captcha, please try again")
)
