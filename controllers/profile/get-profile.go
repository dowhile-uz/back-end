package profilecontrollerfx

import (
	"context"

	usermodelfx "dowhile.uz/back-end/models/user"
	"github.com/danielgtaylor/huma/v2"
)

type (
	GetProfileInput  struct{}
	GetProfileOutput struct {
		Body *usermodelfx.User
	}
)

func (c *Controller) GetProfileHandler(ctx context.Context, input *GetProfileInput) (*GetProfileOutput, error) {
	user, ok := ctx.Value("user").(*usermodelfx.User)
	if !ok {
		return nil, huma.Error401Unauthorized("User is not authorized")
	}

	return &GetProfileOutput{
		Body: user,
	}, nil
}
