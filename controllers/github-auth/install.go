package githubAuthControllerFx

import (
	"context"
	"net/http"
)

type (
	InstallInput  struct{}
	InstallOutput struct {
		Status int
		URL    string `header:"Location"`
	}
)

func (c *Controller) InstallHandler(ctx context.Context, input *InstallInput) (*InstallOutput, error) {
	o := &InstallOutput{
		Status: http.StatusTemporaryRedirect,
		URL:    c.service.GetGithubInstallationURL(),
	}

	return o, nil
}
