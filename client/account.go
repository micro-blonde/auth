package client

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/errors/grpc"
	"github.com/micro-blonde/auth/proto/auth/account"
)

func (c *client) GetAccount(ctx context.Context,
	in *account.GetRequest) (*account.Account, errors.Error) {
	r, err := c.grpcClient.GetAccount(ctx, in)
	if err != nil {
		return nil, grpc.Parse(err).
			WithDesc("error on get account").
			WithTrace("grpcClient.GetAccount.Err")
	}
	return r, nil
}

func (c *client) ListAccounts(ctx context.Context,
	in *account.ListRequest) (*account.Accounts, errors.Error) {
	r, err := c.grpcClient.ListAccounts(ctx, in)
	if err != nil {
		return nil, grpc.Parse(err).
			WithDesc("error on list account").
			WithTrace("grpcClient.ListAccounts.Err")
	}
	return r, nil
}
