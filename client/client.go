package client

import (
	"time"

	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/errors"
	"github.com/ginger-core/log"
	"github.com/micro-blonde/auth/proto/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
)

type Client interface {
	auth.AuthClient
	Initialize()
	IsConnected() bool
}

type client struct {
	logger log.Logger
	config config

	conn        *grpc.ClientConn
	dialOptions []grpc.DialOption
	client      auth.AuthClient
}

func New(logger log.Logger, registry registry.Registry) Client {
	c := &client{
		logger: logger,
		dialOptions: []grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		},
	}
	if err := registry.Unmarshal(&c.config); err != nil {
		panic(err)
	}
	c.config.Initialize()
	return c
}

func (c *client) Initialize() {
	go c.ensureConnection()
}

func (c *client) IsConnected() bool {
	if c.conn == nil {
		return false
	}
	state := c.conn.GetState()
	if state != connectivity.Ready && state != connectivity.Connecting {
		return false
	}
	return true
}

func (c *client) ensureService() errors.Error {
	if !c.config.enabled {
		return errors.Forbidden().
			WithId("AuthClientNotEnabledError").
			WithMessage("client is not enabled to perform your request")
	}
	if c.conn == nil {
		return errors.Internal().
			WithId("ServiceUnavailable").
			WithMessage("Service is unavailable right now")
	}
	return nil
}

func (c *client) ensureConnection() {
	if !c.config.enabled {
		return
	}
	for {
		if c.conn != nil {
			time.Sleep(time.Second)
			state := c.conn.GetState()
			if state == connectivity.Ready || state == connectivity.Connecting {
				continue
			}
		}
		c.logger.Infof("connecting to server to address %s...", c.config.Address)
		var err error
		c.conn, err = grpc.NewClient(c.config.Address, c.dialOptions...)
		if err != nil {
			c.logger.Infof("connection could not be established, retrying...")
			time.Sleep(time.Second)
			continue
		}
		c.conn.Connect()
		c.logger.Infof("connected. state: %s", c.conn.GetState().String())
		c.client = auth.NewAuthClient(c.conn)
	}
}
