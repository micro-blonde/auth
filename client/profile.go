package client

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/errors/grpc"
	"github.com/micro-blonde/auth/proto/auth/account/profile"
)

func (c *client) GetProfile(ctx context.Context,
	in *profile.GetRequest) (*profile.Profile, errors.Error) {
	r, err := c.grpcClient.GetProfile(ctx, in)
	if err != nil {
		return nil, grpc.Parse(err).
			WithDesc("error on get profile").
			WithTrace("grpcClient.GetAccount.Err")
	}
	return r, nil
}
