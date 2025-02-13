package editorControllerFx

import (
	"net/http"

	authorizedMiddlewareLibFx "dowhile.uz/back-end/lib/authorized-middleware"
	editorServiceFx "dowhile.uz/back-end/services/editor"
	"dowhile.uz/back-end/utils"
	"github.com/danielgtaylor/huma/v2"
	"go.uber.org/fx"
)

var Module = fx.Module("controllers.editor", fx.Provide(New))

type (
	Params struct {
		fx.In
		Service              editorServiceFx.Service
		AuthorizedMiddleware *authorizedMiddlewareLibFx.Middleware
	}
	Controller struct {
		service              editorServiceFx.Service
		authorizedMiddleware *authorizedMiddlewareLibFx.Middleware
	}
)

func (c *Controller) Routes(api huma.API) {
	huma.Register(api, huma.Operation{
		Path:        "/v1/editor/user-repo-availabitily-check",
		Method:      http.MethodPost,
		OperationID: "v1-editor-user-repo-availability-check",
		Tags:        []string{"v1", "protected", "editor", "editor-project"},
		Summary:     "Check for repository name availability",
		Security:    utils.JWTSecurityScheme,
		Middlewares: huma.Middlewares{
			c.authorizedMiddleware.GetMiddleware(api),
		},
	}, c.UserRepoAvailabilityCheckHandler)

	huma.Register(api, huma.Operation{
		Path:        "/v1/editor/community-repo-availabitily-check",
		Method:      http.MethodPost,
		OperationID: "v1-editor-community-repo-availability-check",
		Tags:        []string{"v1", "protected", "editor", "editor-project"},
		Summary:     "Check for repository name availability",
		Security:    utils.JWTSecurityScheme,
		Middlewares: huma.Middlewares{
			c.authorizedMiddleware.GetMiddleware(api),
		},
	}, c.CommunityRepoAvailabilityCheckHandler)

	huma.Register(api, huma.Operation{
		Path:        "/v1/editor/project",
		Method:      http.MethodPost,
		OperationID: "v1-editor-project-create",
		Tags:        []string{"v1", "protected", "editor", "editor-project"},
		Summary:     "Create new editor project",
		Security:    utils.JWTSecurityScheme,
		Middlewares: huma.Middlewares{
			c.authorizedMiddleware.GetMiddleware(api),
		},
	}, c.CreateProjectHandler)
}

func New(p Params) Controller {
	return Controller{
		service:              p.Service,
		authorizedMiddleware: p.AuthorizedMiddleware,
	}
}
