package grpc

import (
	"context"

	"github.com/micro-blonde/auth/proto/auth/account"
	"google.golang.org/grpc"
)

func (c *client) GetAccount(ctx context.Context,
	in *account.GetRequest, opts ...grpc.CallOption) (*account.Account, error) {
	result := new(account.Account)
	if err := c.ensureService(); err != nil {
		return result, err
	}
	result, err := c.client.GetAccount(ctx, in, opts...)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (c *client) ListAccounts(ctx context.Context,
	in *account.ListRequest, opts ...grpc.CallOption) (*account.Accounts, error) {
	result := new(account.Accounts)
	if err := c.ensureService(); err != nil {
		return result, err
	}
	result, err := c.client.ListAccounts(ctx, in, opts...)
	if err != nil {
		return result, err
	}
	return result, nil
}
