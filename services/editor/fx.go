package editorServiceFx

import "go.uber.org/fx"

var Module = fx.Module("services.editor", fx.Provide(New))

type (
	Params struct {
		fx.In
	}
	Service struct{}
)

func New(p Params) Service {
	return Service{}
}
