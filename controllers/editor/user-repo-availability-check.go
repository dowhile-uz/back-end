package editorControllerFx

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
)

type (
	UserRepoAvailabilityCheckInput struct {
		Body struct {
			Owner    string `json:"owner"`
			RepoName string `json:"repo-name"`
		}
	}
	UserRepoAvailabilityCheckOutput struct {
		Body struct {
			IsAvailable bool `json:"is-available"`
		}
	}
)

func (c *Controller) UserRepoAvailabilityCheckHandler(ctx context.Context, input *UserRepoAvailabilityCheckInput) (*UserRepoAvailabilityCheckOutput, error) {
	accessToken, ok := ctx.Value("access_token").(string)
	if !ok {
		return nil, huma.Error401Unauthorized("User not authorized")
	}

	isAvailable, err := c.service.UserRepoAvailabilityCheck(ctx, accessToken, input.Body.Owner, input.Body.RepoName)
	if err != nil {
		return nil, err
	}

	output := &UserRepoAvailabilityCheckOutput{}
	output.Body.IsAvailable = isAvailable

	return output, nil
}
