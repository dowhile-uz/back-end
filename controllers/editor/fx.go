package editorControllerFx

import (
	editorServiceFx "dowhile.uz/back-end/services/editor"
	"github.com/danielgtaylor/huma/v2"
	"go.uber.org/fx"
)

var Module = fx.Module("controllers.editor", fx.Provide(New))

type (
	Params struct {
		fx.In
		Service editorServiceFx.Service
	}
	Controller struct {
		service editorServiceFx.Service
	}
)

func (c *Controller) Routes(api huma.API) {
	// huma.Register(api, huma.Operation{
	// 	Path:        "/v1/github-auth/redirect",
	// 	Method:      http.MethodGet,
	// 	OperationID: "v1-github-auth-redirect",
	// 	Tags:        []string{"v1", "public", "auth", "github-auth"},
	// 	Summary:     "Redirect to the GitHub authotization page",
	// }, c.RedirectHandler)
}

func New(p Params) Controller {
	return Controller{
		service: p.Service,
	}
}
