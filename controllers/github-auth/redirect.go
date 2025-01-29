package githubauthcontrollerfx

import (
	"context"
)

type (
	RedirectInput struct {
		Back string `query:"back"`
	}
	RedirectOutput struct {
		Body struct {
			RedirectURL string `json:"redirectUrl" example:"https://github.com/login/oauth/authorize" doc:"GitHub Auth Redirect URL"`
		}
	}
)

func (c *Controller) RedirectHandler(ctx context.Context, input *RedirectInput) (*RedirectOutput, error) {
	o := &RedirectOutput{}
	o.Body.RedirectURL = c.Service.GetRedirectURL(input.Back)
	return o, nil
}
