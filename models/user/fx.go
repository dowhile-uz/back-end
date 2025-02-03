package usermodelfx

import (
	postgreslibfx "dowhile.uz/back-end/lib/postgres"
	redislibfx "dowhile.uz/back-end/lib/redis"
	"go.uber.org/fx"
)

var Module = fx.Module("models.user", fx.Provide(New))

type (
	Params struct {
		fx.In

		Postgres *postgreslibfx.Postgres
		Redis    *redislibfx.Redis
	}
	Model struct {
		postgres *postgreslibfx.Postgres
		redis    *redislibfx.Redis
	}
)

func New(p Params) *Model {
	return &Model{
		postgres: p.Postgres,
		redis:    p.Redis,
	}
}
