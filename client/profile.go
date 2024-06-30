package client

import (
	"context"

	"github.com/micro-blonde/auth/proto/auth/account/profile"
	"google.golang.org/grpc"
)

func (c *client) GetProfile(ctx context.Context, in *profile.GetRequest,
	opts ...grpc.CallOption) (*profile.Profile, error) {
	result := new(profile.Profile)
	if err := c.ensureService(); err != nil {
		return result, err
	}
	result, err := c.client.GetProfile(ctx, in, opts...)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (c *client) ListProfiles(ctx context.Context,
	in *profile.ListRequest, opts ...grpc.CallOption) (*profile.Profiles, error) {
	result := new(profile.Profiles)
	if err := c.ensureService(); err != nil {
		return result, err
	}
	result, err := c.client.ListProfiles(ctx, in, opts...)
	if err != nil {
		return result, err
	}
	return result, nil
}
