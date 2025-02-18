package githubControllerFx

import (
	"net/http"

	authorizedMiddlewareLibFx "dowhile.uz/back-end/lib/authorized-middleware"
	githubServiceFx "dowhile.uz/back-end/services/github"
	"dowhile.uz/back-end/utils"
	"github.com/danielgtaylor/huma/v2"
	"go.uber.org/fx"
)

var Module = fx.Module("controllers.github", fx.Provide(New))

type (
	Params struct {
		fx.In
		AuthorizedMiddleware *authorizedMiddlewareLibFx.Middleware
		Service              githubServiceFx.Service
	}
	Controller struct {
		authorizedMiddleware *authorizedMiddlewareLibFx.Middleware
		service              githubServiceFx.Service
	}
)

func (c *Controller) Routes(api huma.API) {
	huma.Register(api, huma.Operation{
		Path:        "/v1/github/list-repos",
		Method:      http.MethodGet,
		OperationID: "v1-github-list-repos",
		Tags:        []string{"v1", "protected", "github"},
		Summary:     "List all installed repos",
		Security:    utils.JWTSecurityScheme,
		Middlewares: huma.Middlewares{
			c.authorizedMiddleware.GetMiddleware(api),
		},
	}, c.ListInstallationRepos)
}

func New(p Params) Controller {
	return Controller{
		authorizedMiddleware: p.AuthorizedMiddleware,
		service:              p.Service,
	}
}
