package authorization

import (
	"time"
)

const (
	CaptchaErrorCode = 'C'
)

type captcha struct {
	Debug      bool
	Expiration time.Duration
	Skew       *uint
	Digits     uint
}

func (c *captcha) initialize() {
	if c.Expiration == 0 {
		c.Expiration = time.Second * 30
	}
	if c.Digits == 0 {
		c.Digits = 6
	}
}

type config struct {
	AccessKeyPrefix string
	Captcha         captcha
}

func (c *config) initialize() {
	if c.AccessKeyPrefix == "" {
		c.AccessKeyPrefix = "access_"
	}
	c.Captcha.initialize()
}
