package errorCodesLibFx

import "go.uber.org/fx"

var Module = fx.Module("lib.error-codes", fx.Provide(New))

type Lib struct{}

func New() *Lib {
	return &Lib{}
}
