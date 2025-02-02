package main

import (
	"fmt"

	configlibfx "dowhile.uz/back-end/lib/config"
)

func main() {
	config, err := configlibfx.New(configlibfx.Params{})
	if err != nil {
		panic(err)
	}

	fmt.Printf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.Infrastructure.Postgres.User,
		config.Infrastructure.Postgres.Password,
		config.Infrastructure.Postgres.Host,
		config.Infrastructure.Postgres.Port,
		config.Infrastructure.Postgres.DB,
	)
}
