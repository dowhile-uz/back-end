package helloworldcontrollerfx

import (
	"github.com/danielgtaylor/huma/v2"
	"go.uber.org/fx"
)

var Module = fx.Module("controllers.hello-world", fx.Provide(New))

type (
	Params struct {
		fx.In
	}
	HelloWorldController struct{}
)

func (h *HelloWorldController) Routes(api huma.API) {
	huma.Get(api, "/", h.HelloWorldHandler)
}

func New(_ Params) HelloWorldController {
	return HelloWorldController{}
}
