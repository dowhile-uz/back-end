package main

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"dowhile.uz/back-end/lib/router"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		router.Module,
		fx.Invoke(New),
	).Run()
}

func New(lc fx.Lifecycle, router http.Handler) {
	server := http.Server{
		Addr:    "0.0.0.0:8000",
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
