package helloworld

import (
	"github.com/danielgtaylor/huma/v2"
	"go.uber.org/fx"
)

var Module = fx.Module("services.hello-world", fx.Provide(New))

type (
	Options struct {
		fx.In
	}
	HelloWorldService struct{}
)

func (h *HelloWorldService) Routes(api huma.API) {
	huma.Get(api, "/", h.HelloWorldHandler)
}

func New() HelloWorldService {
	return HelloWorldService{}
}
