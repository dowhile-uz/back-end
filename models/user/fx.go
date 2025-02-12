package userModelFx

import (
	postgresLibFx "dowhile.uz/back-end/lib/postgres"
	redisLibFx "dowhile.uz/back-end/lib/redis"
	"go.uber.org/fx"
)

var Module = fx.Module("models.user", fx.Provide(New))

type (
	Params struct {
		fx.In

		Postgres *postgresLibFx.Postgres
		Redis    *redisLibFx.Redis
	}
	Model struct {
		postgres *postgresLibFx.Postgres
		redis    *redisLibFx.Redis
	}
)

func New(p Params) *Model {
	return &Model{
		postgres: p.Postgres,
		redis:    p.Redis,
	}
}
