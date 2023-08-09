package main

import (
	"context"
	"fmt"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net"
	"net/http"
)

func NewHTTPServer(lc fx.Lifecycle, mux *http.ServeMux, param Params) *http.Server {
	srv := &http.Server{Addr: ":8080", Handler: mux}

	log := param.Logger
	fmt.Print("test1\n")
	if log == nil {
		fmt.Print("test2\n")
		log = zap.NewNop()
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			log.Info("Starting HTTP server", zap.String("addr", srv.Addr))
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
