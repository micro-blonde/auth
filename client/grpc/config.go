package grpc

import (
	"time"
)

type config struct {
	Address string
	Timeout time.Duration
}

func (c *config) Initialize() {
	if c.Timeout == 0 {
		c.Timeout = time.Second * 3
	}
}
