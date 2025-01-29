package main

import (
	"context"
	"fmt"
	"net"
	"net/http"

	controllersfx "dowhile.uz/back-end/controllers"
	libfx "dowhile.uz/back-end/lib"
	configlibfx "dowhile.uz/back-end/lib/config"
	servicesfx "dowhile.uz/back-end/services"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		libfx.Module,
		servicesfx.Module,
		controllersfx.Module,
		fx.Invoke(New),
	).Run()
}

func New(lc fx.Lifecycle, router http.Handler, config *configlibfx.Config) {
	server := http.Server{
		Addr:    fmt.Sprintf("%s:%v", config.Server.Host, config.Server.Port),
		Handler: router,
	}

	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			ln, err := net.Listen("tcp", server.Addr)
			if err != nil {
				return err
			}

			fmt.Println("Starting HTTP server at", server.Addr)

			go server.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})
}
