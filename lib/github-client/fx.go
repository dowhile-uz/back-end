package githubclientlibfx

import (
	configlibfx "dowhile.uz/back-end/lib/config"
	"github.com/google/go-github/v68/github"
	"go.uber.org/fx"
)

var Module = fx.Module("lib.github-client", fx.Provide(New))

type (
	Params struct {
		fx.In
		Config *configlibfx.Config
	}
	Client struct {
		github.Client
	}
)

func New(p Params) *Client {
	return &Client{
		Client: *github.NewClient(nil),
	}
}
