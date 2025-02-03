package profilecontrollerfx

import (
	"net/http"

	authorizedmiddlewarelibfx "dowhile.uz/back-end/lib/authorized-middleware"
	"dowhile.uz/back-end/utils"
	"github.com/danielgtaylor/huma/v2"
	"go.uber.org/fx"
)

var Module = fx.Module("controllers.profile", fx.Provide(New))

type (
	Params struct {
		fx.In
		AuthorizedMiddleware *authorizedmiddlewarelibfx.Middleware
	}
	Controller struct {
		authorizedMiddleware *authorizedmiddlewarelibfx.Middleware
	}
)

func (c *Controller) Routes(api huma.API) {
	huma.Register(api, huma.Operation{
		Path:        "/v1/profile",
		Method:      http.MethodGet,
		OperationID: "v1-profile",
		Tags:        []string{"v1", "protected", "auth", "user"},
		Summary:     "Get current user profile",
		Security:    utils.JWTSecurityScheme,
		Middlewares: huma.Middlewares{
			c.authorizedMiddleware.GetMiddleware(api),
		},
	}, c.GetProfileHandler)
}

func New(p Params) Controller {
	return Controller{
		authorizedMiddleware: p.AuthorizedMiddleware,
	}
}
