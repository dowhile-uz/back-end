package usermodelfx

import (
	postgreslibfx "dowhile.uz/back-end/lib/postgres"
	"go.uber.org/fx"
)

var Module = fx.Module("models.user", fx.Provide(New))

type (
	Params struct {
		fx.In

		Postgres *postgreslibfx.Postgres
	}
	Model struct {
		Postgres *postgreslibfx.Postgres
	}
)

func New(p Params) *Model {
	return &Model{
		Postgres: p.Postgres,
	}
}
