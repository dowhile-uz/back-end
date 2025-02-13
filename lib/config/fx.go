package configLibFx

import (
	"os"

	"github.com/danielgtaylor/huma/v2"
	"go.uber.org/config"
	"go.uber.org/fx"
)

var Module = fx.Module("lib.config", fx.Provide(New))

type (
	Params struct {
		fx.In
	}
	Config struct {
		Server struct {
			Host        string
			Port        string
			PublicURL   string `yaml:"public-url"`
			FrontEndURL string `yaml:"front-end-url"`
			Cors        struct {
				AllowedOrigins []string `yaml:"allowed-origins"`
			}
			JWTSecret string `yaml:"jwt-secret"`
		}

		Github struct {
			BotPersonalToken string `yaml:"bot-personal-token"`
			CommunityOrgName string `yaml:"community-org-name"`
		}

		GithubAuth struct {
			AppID                int64  `yaml:"app-id"`
			ClientID             string `yaml:"client-id"`
			ClientSecret         string `yaml:"client-secret"`
			RedirectCompletePath string `yaml:"redirect-complete-path"`
			RedirectFrontEndPath string `yaml:"redirect-front-end-path"`
			Scopes               []string
		} `yaml:"github-auth"`

		Infrastructure struct {
			Postgres struct {
				Host     string
				Port     int
				DB       string `yaml:"db"`
				User     string
				Password string
			}
			Redis struct {
				Host     string
				Port     int
				Password string
			}
		}

		OpenAPI huma.OpenAPI `yaml:"openapi"`
	}
)

func New(p Params) (*Config, error) {
	c := &Config{}

	base, err := os.Open("configs/base.yaml")
	if err != nil {
		return nil, err
	}

	var provider config.Provider

	override, err := os.Open("configs/override.yaml")

	if err == nil {
		provider, err = config.NewYAML(config.Source(base), config.Source(override))
		override.Close()
	} else {
		provider, err = config.NewYAML(config.Source(base))
	}

	base.Close()

	if err != nil {
		return nil, err
	}

	err = provider.Get(config.Root).Populate(&c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
