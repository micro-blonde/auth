package client

import (
	"time"
)

type config struct {
	Enabled *bool
	Address string
	Timeout time.Duration

	enabled bool
}

func (c *config) Initialize() {
	if c.Timeout == 0 {
		c.Timeout = time.Second * 3
	}
	c.enabled = c.Enabled == nil || *c.Enabled
}
