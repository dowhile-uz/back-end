package githubauthservicefx

import "go.uber.org/fx"

var Module = fx.Module("services.github-auth", fx.Provide(New))

type (
	Params struct {
		fx.In
	}
	GithubAuthService struct{}
)

func New(_ Params) GithubAuthService {
	return GithubAuthService{}
}
