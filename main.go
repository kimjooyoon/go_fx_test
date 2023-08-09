package main

import (
	"context"
	"go.uber.org/fx"
	"net"
	"net/http"
)

func main() {
	fx.New(
		fx.Provide(NewHttpServer),
	).Run()
}

func NewHttpServer(lc fx.Lifecycle) *http.Server {
	srv := &http.Server{Addr: ":8080"}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			go func() {
				err := srv.Serve(ln)
				if err != nil {
					return
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})

	return srv
}
