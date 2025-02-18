package githubServiceFx

import (
	configLibFx "dowhile.uz/back-end/lib/config"
	githubClientLibFx "dowhile.uz/back-end/lib/github-client"
	userModelFx "dowhile.uz/back-end/models/user"
	"go.uber.org/fx"
)

var Module = fx.Module("services.github", fx.Provide(New))

type (
	Params struct {
		fx.In
		GithubClient *githubClientLibFx.Client
		Config       *configLibFx.Config
		UserModel    *userModelFx.Model
	}
	Service struct {
		githubClient *githubClientLibFx.Client
		config       *configLibFx.Config
		userModel    *userModelFx.Model
	}
)

func New(p Params) Service {
	return Service{
		githubClient: p.GithubClient,
		config:       p.Config,
		userModel:    p.UserModel,
	}
}
