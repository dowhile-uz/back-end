package routerlibfx

import (
	"net/http"

	helloworldcontrollerfx "dowhile.uz/back-end/controllers/hello-world"
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
)

var Module = fx.Module("lib.router", fx.Provide(New), helloworldcontrollerfx.Module)

type Params struct {
	fx.In
	HelloWorld helloworldcontrollerfx.HelloWorldController
}

func New(p Params) http.Handler {
	router := chi.NewMux()

	api := humachi.New(router, huma.DefaultConfig("dowhile.uz API", "1.0.0"))

	p.HelloWorld.Routes(api)

	return router
}
