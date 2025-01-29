package configlibfx

import (
	"fmt"
	"os"

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
			Host string
			Port string
		}
		Githubauth struct {
			Clientid string
		}
	}
)

func New(p Params) (*Config, error) {
	c := &Config{}

	file, err := os.Open("configs/base.yaml")
	if err != nil {
		return nil, err
	}

	provider, err := config.NewYAML(config.Source(file))
	if err != nil {
		return nil, err
	}

	err = provider.Get(config.Root).Populate(&c)
	if err != nil {
		return nil, err
	}

	fmt.Println("server", provider.Get(config.Root), c)

	return c, nil
}
