package githubauthservicefx

import (
	configlibfx "dowhile.uz/back-end/lib/config"
	githubclientlibfx "dowhile.uz/back-end/lib/github-client"
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
	}
	Service struct {
		GithubClient *githubclientlibfx.Client
		Config       *configlibfx.Config
	}
)

func New(p Params) Service {
	return Service{
		GithubClient: p.GithubClient,
		Config:       p.Config,
	}
}
