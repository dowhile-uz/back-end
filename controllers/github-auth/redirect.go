package githubauthcontrollerfx

import (
	"context"
	"net/http"
)

type (
	RedirectInput struct {
		Back string `query:"back"`
	}
	RedirectOutput struct {
		Status int
		URL    string `header:"Location"`
	}
)

func (c *Controller) RedirectHandler(ctx context.Context, input *RedirectInput) (*RedirectOutput, error) {
	o := &RedirectOutput{
		Status: http.StatusTemporaryRedirect,
		URL:    c.Service.GetGitHubRedirectURL(input.Back),
	}

	return o, nil
}
