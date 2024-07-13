package captcha

import (
	"time"
)

type Captcha struct {
	Image      Image
	Secret     string
	Expiration time.Duration
	Code       string
}
