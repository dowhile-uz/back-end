package redisLibFx

import (
	"context"
	"fmt"
	"time"

	configLibFx "dowhile.uz/back-end/lib/config"
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
)

var Module = fx.Module("lib.redis", fx.Provide(New))

type (
	Params struct {
		fx.In
		Config *configLibFx.Config
	}
	Redis struct {
		redis.Client
	}
)

func New(p Params) (*Redis, error) {
	c := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", p.Config.Infrastructure.Redis.Host, p.Config.Infrastructure.Redis.Port),
		Password: p.Config.Infrastructure.Redis.Password,
		DB:       0,
	})

	status := c.Set(context.Background(), "_initialized_", true, time.Duration(1))
	if status != nil && status.Err() != nil {
		return nil, status.Err()
	}

	return &Redis{
		Client: *c,
	}, nil
}
