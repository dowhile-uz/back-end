package githubauthcontrollerfx

import (
	configlibfx "dowhile.uz/back-end/lib/config"
	usermodelfx "dowhile.uz/back-end/models/user"
	githubauthservicefx "dowhile.uz/back-end/services/github-auth"
	"github.com/danielgtaylor/huma/v2"
	"go.uber.org/fx"
)

var Module = fx.Module("controllers.github-auth", fx.Provide(New))

type (
	Params struct {
		fx.In
		Service   githubauthservicefx.Service
		UserModel *usermodelfx.Model
		Config    *configlibfx.Config
	}
	Controller struct {
		Service   githubauthservicefx.Service
		UserModel *usermodelfx.Model
		Config    *configlibfx.Config
	}
)

func (c *Controller) Routes(api huma.API) {
	huma.Get(api, "/v1/github-auth/redirect", c.RedirectHandler)
	huma.Get(api, "/v1/github-auth/complete", c.CompleteHandler)
}

func New(p Params) Controller {
	return Controller{
		Service:   p.Service,
		UserModel: p.UserModel,
		Config:    p.Config,
	}
}
