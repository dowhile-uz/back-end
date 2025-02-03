package githubauthservicefx

import (
	configlibfx "dowhile.uz/back-end/lib/config"
	githubclientlibfx "dowhile.uz/back-end/lib/github-client"
	postgreslibfx "dowhile.uz/back-end/lib/postgres"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"services.github-auth",
	fx.Provide(New),
)

type (
	Params struct {
		fx.In
		GithubClient *githubclientlibfx.Client
		Config       *configlibfx.Config
		Postgres     *postgreslibfx.Postgres
	}
	Service struct {
		githubClient *githubclientlibfx.Client
		config       *configlibfx.Config
	}
)

func New(p Params) Service {
	return Service{
		githubClient: p.GithubClient,
		config:       p.Config,
	}
}
