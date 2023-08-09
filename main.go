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
		fx.Invoke(func(server2 Server) {}),
	).Run()
}

type Server interface {
	getAddr() string
	runServe(l net.Listener) error
	Shutdown(ctx context.Context) error
}

type api struct {
	*http.Server
}

func (r api) getAddr() string               { return r.Addr }
func (r api) runServe(l net.Listener) error { return r.Serve(l) }

func NewHttpServer(lc fx.Lifecycle) Server {
	var srv Server = api{&http.Server{Addr: ":8080"}}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.getAddr())
			if err != nil {
				return err
			}
			go func() {
				err := srv.runServe(ln)
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
