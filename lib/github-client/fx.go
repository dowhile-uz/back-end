package githubClientLibFx

import (
	"context"
	"os"

	configLibFx "dowhile.uz/back-end/lib/config"
	"github.com/google/go-github/v68/github"
	"github.com/jferrl/go-githubauth"
	"go.uber.org/fx"
	"golang.org/x/oauth2"
)

var Module = fx.Module("lib.github-client", fx.Provide(New))

type (
	Params struct {
		fx.In
		Config *configLibFx.Config
	}
	Client struct {
		*github.Client
		Bot                    github.Client
		applicationTokenSource oauth2.TokenSource
	}
)

func (c *Client) InstallationClient(installationID int64) *github.Client {
	installationTokenSource := githubauth.NewInstallationTokenSource(installationID, c.applicationTokenSource)
	httpClient := oauth2.NewClient(context.Background(), installationTokenSource)
	return github.NewClient(httpClient)
}

func New(p Params) (*Client, error) {
	client := &Client{
		Client: github.NewClient(nil),
	}

	privateKey, err := os.ReadFile("configs/github-app.pem")
	if err != nil {
		return nil, err
	}

	applicationTokenSource, err := githubauth.NewApplicationTokenSource(p.Config.GithubAuth.AppID, privateKey)
	if err != nil {
		return nil, err
	}

	client.applicationTokenSource = applicationTokenSource

	return client, nil
}
