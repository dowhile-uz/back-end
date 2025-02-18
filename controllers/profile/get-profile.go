package profileControllerFx

import (
	"context"

	userModelFx "dowhile.uz/back-end/models/user"
	"github.com/danielgtaylor/huma/v2"
)

type (
	InternalUser struct {
		*userModelFx.User
	}
	GetProfileInput  struct{}
	GetProfileOutput struct {
		Body InternalUser
	}
)

func (c *Controller) GetProfileHandler(ctx context.Context, input *GetProfileInput) (*GetProfileOutput, error) {
	user, ok := ctx.Value("user").(*userModelFx.User)
	if !ok {
		return nil, huma.Error401Unauthorized("User is not authorized")
	}

	output := &GetProfileOutput{}
	output.Body.User = user

	return output, nil
}
