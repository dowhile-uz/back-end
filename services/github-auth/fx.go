package githubAuthServiceFx

import (
	configLibFx "dowhile.uz/back-end/lib/config"
	githubClientLibFx "dowhile.uz/back-end/lib/github-client"
	postgresLibFx "dowhile.uz/back-end/lib/postgres"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"services.github-auth",
	fx.Provide(New),
)

type (
	Params struct {
		fx.In
		GithubClient *githubClientLibFx.Client
		Config       *configLibFx.Config
		Postgres     *postgresLibFx.Postgres
	}
	Service struct {
		githubClient *githubClientLibFx.Client
		config       *configLibFx.Config
	}
)

func New(p Params) Service {
	return Service{
		githubClient: p.GithubClient,
		config:       p.Config,
	}
}
