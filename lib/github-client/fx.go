package githubClientLibFx

import (
	configLibFx "dowhile.uz/back-end/lib/config"
	"github.com/google/go-github/v68/github"
	"go.uber.org/fx"
)

var Module = fx.Module("lib.github-client", fx.Provide(New))

type (
	Params struct {
		fx.In
		Config *configLibFx.Config
	}
	Client struct {
		github.Client
		Bot github.Client
	}
)

func New(p Params) *Client {
	// basicAuthTransport := github.BasicAuthTransport{
	// 	Username: p.Config.GithubAuth.ClientID,
	// 	Password: p.Config.GithubAuth.ClientSecret,
	// }

	client := &Client{
		// Client: *github.NewClient(basicAuthTransport.Client()),

		Client: *github.NewClient(nil),
		Bot:    *github.NewClient(nil).WithAuthToken(p.Config.Github.BotPersonalToken),
	}

	return client
}
