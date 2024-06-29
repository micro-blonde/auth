package client

import (
	"context"

	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/errors"
	"github.com/ginger-core/log"
	"github.com/micro-blonde/auth/client/grpc"
	aproto "github.com/micro-blonde/auth/proto/auth/account"
	"github.com/micro-blonde/auth/proto/auth/account/profile"
)

type Client interface {
	Initialize()
	GetGrpcClient() grpc.Client

	GetAccount(ctx context.Context,
		in *aproto.GetRequest) (*aproto.Account, errors.Error)
	GetProfile(ctx context.Context,
		in *profile.GetRequest) (*profile.Profile, errors.Error)

	ListAccounts(ctx context.Context,
		in *aproto.ListRequest) (*aproto.Accounts, errors.Error)
}

type client struct {
	logger log.Logger

	grpcClient grpc.Client
}

func NewClient(logger log.Logger, registry registry.Registry) Client {
	c := &client{
		logger:     logger,
		grpcClient: grpc.New(logger, registry),
	}
	return c
}

func (c *client) GetGrpcClient() grpc.Client {
	return c.grpcClient
}

func (c *client) Initialize() {
	c.grpcClient.Initialize()
}
