package router

import (
	"net/http"

	helloworld "dowhile.uz/back-end/services/hello-world"
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
)

var Module = fx.Module("lib.router", fx.Provide(New), helloworld.Module)

type Options struct {
	fx.In
	HelloWorld helloworld.HelloWorldService
}

func New(opts Options) http.Handler {
	router := chi.NewMux()

	api := humachi.New(router, huma.DefaultConfig("dowhile.uz API", "1.0.0"))

	opts.HelloWorld.Routes(api)

	return router
}
