package routerlibfx

import (
	"net/http"

	githubauthcontrollerfx "dowhile.uz/back-end/controllers/github-auth"
	helloworldcontrollerfx "dowhile.uz/back-end/controllers/hello-world"
	configlibfx "dowhile.uz/back-end/lib/config"
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
)

var Module = fx.Module("lib.router", fx.Provide(New))

type Params struct {
	fx.In
	Config     *configlibfx.Config
	HelloWorld helloworldcontrollerfx.Controller
	GithubAuth githubauthcontrollerfx.Controller
}

func New(p Params) http.Handler {
	router := chi.NewMux()

	config := huma.DefaultConfig("dowhile.uz", "1.0.0")
	config.Servers = p.Config.Openapi.Servers

	api := humachi.New(router, config)

	p.HelloWorld.Routes(api)
	p.GithubAuth.Routes(api)

	return router
}
