package postgreslibfx

import (
	"fmt"

	configlibfx "dowhile.uz/back-end/lib/config"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/fx"
)

var Module = fx.Module("lib.postgres", fx.Provide(New))

type (
	Params struct {
		fx.In

		Config *configlibfx.Config
	}
	Postgres struct {
		*sqlx.DB
	}
)

func New(p Params) (*Postgres, error) {
	db, err := sqlx.Connect(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			p.Config.Infrastructure.Postgres.Host,
			p.Config.Infrastructure.Postgres.Port,
			p.Config.Infrastructure.Postgres.User,
			p.Config.Infrastructure.Postgres.Password,
			p.Config.Infrastructure.Postgres.DB,
		),
	)
	if err != nil {
		return nil, err
	}

	return &Postgres{DB: db}, nil
}
