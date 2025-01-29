package configlibfx

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
			Host      string
			Port      string
			Publicurl string
		}
		Githubauth struct {
			Appid        int64
			Clientid     string
			Clientsecret string
			Redirectpath string
		}
		Openapi huma.OpenAPI
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
