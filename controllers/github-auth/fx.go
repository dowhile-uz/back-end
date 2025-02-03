package githubauthcontrollerfx

import (
	"net/http"

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
		service   githubauthservicefx.Service
		userModel *usermodelfx.Model
		config    *configlibfx.Config
	}
)

func (c *Controller) Routes(api huma.API) {
	huma.Register(api, huma.Operation{
		Path:        "/v1/github-auth/redirect",
		Method:      http.MethodGet,
		OperationID: "v1-github-auth-redirect",
		Tags:        []string{"v1", "public", "auth", "github-auth"},
		Summary:     "Redirect to the GitHub authotization page",
	}, c.RedirectHandler)

	huma.Register(api, huma.Operation{
		Path:        "/v1/github-auth/complete",
		Method:      http.MethodGet,
		OperationID: "v1-github-auth-complete",
		Tags:        []string{"v1", "public", "auth", "github-auth"},
		Summary:     "Handle code from the GitHub authorization page",
	}, c.CompleteHandler)
}

func New(p Params) Controller {
	return Controller{
		service:   p.Service,
		userModel: p.UserModel,
		config:    p.Config,
	}
}
