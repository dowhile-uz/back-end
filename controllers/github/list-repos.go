package githubControllerFx

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"github.com/google/go-github/v68/github"
)

type (
	ListInstallationReposInput  struct{}
	ListInstallationReposOutput struct {
		Body []*github.Repository
	}
)

func (c *Controller) ListInstallationRepos(ctx context.Context, input *ListInstallationReposInput) (*ListInstallationReposOutput, error) {
	accessToken, ok := ctx.Value("access_token").(string)
	if !ok {
		return nil, huma.Error401Unauthorized("User is not authorized")
	}

	repos, err := c.service.ListInstallationRepos(ctx, accessToken)
	if err != nil {
		return nil, err
	}

	output := &ListInstallationReposOutput{}
	output.Body = repos
	return output, nil
}
