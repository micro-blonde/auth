package grpc

import (
	"context"

	"github.com/micro-blonde/auth/proto/auth/account"
	"google.golang.org/grpc"
)

func (c *client) ListAccountProfiles(ctx context.Context,
	in *account.ListRequest,
	opts ...grpc.CallOption) (*account.AccountProfiles, error) {
	result := new(account.AccountProfiles)
	if err := c.ensureService(); err != nil {
		return result, err
	}
	result, err := c.client.ListAccountProfiles(ctx, in, opts...)
	if err != nil {
		return result, err
	}
	return result, nil
}
