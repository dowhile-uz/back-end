package routerlibfx

import (
	"net/http"

	githubauthcontrollerfx "dowhile.uz/back-end/controllers/github-auth"
	profilecontrollerfx "dowhile.uz/back-end/controllers/profile"
	configlibfx "dowhile.uz/back-end/lib/config"
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
	"go.uber.org/fx"
)

var Module = fx.Module("lib.router", fx.Provide(New))

type Params struct {
	fx.In
	Config     *configlibfx.Config
	GithubAuth githubauthcontrollerfx.Controller
	Profile    profilecontrollerfx.Controller
}

func New(p Params) http.Handler {
	router := chi.NewMux()

	corsOptions := cors.New(cors.Options{
		AllowedOrigins: p.Config.Server.Cors.AllowedOrigins,
	})

	router.Use(corsOptions.Handler)

	config := huma.DefaultConfig("dowhile.uz", "1.0.0")
	config.Servers = p.Config.OpenAPI.Servers

	config.Components.SecuritySchemes = map[string]*huma.SecurityScheme{
		"access_token": {
			Name:        "Authorization",
			Type:        "apiKey",
			Description: "JWT access token",
			In:          "header",
			Scheme:      "Bearer",
		},
	}

	api := humachi.New(router, config)

	p.GithubAuth.Routes(api)
	p.Profile.Routes(api)

	return router
}
