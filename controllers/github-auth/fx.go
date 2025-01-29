package githubauthcontrollerfx

import (
	githubauthservicefx "dowhile.uz/back-end/services/github-auth"
	"github.com/danielgtaylor/huma/v2"
	"go.uber.org/fx"
)

var Module = fx.Module("controllers.github-auth", fx.Provide(New))

type (
	Params struct {
		fx.In
		Service githubauthservicefx.Service
	}
	Controller struct {
		Service githubauthservicefx.Service
	}
)

func (c *Controller) Routes(api huma.API) {
	huma.Get(api, "/v1/github-auth/redirect", c.RedirectHandler)
	huma.Get(api, "/v1/github-auth/complete", c.CompleteHandler)
}

func New(p Params) Controller {
	return Controller{
		Service: p.Service,
	}
}
