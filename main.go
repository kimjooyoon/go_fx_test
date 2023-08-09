package main

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net/http"
)

type Route interface {
	http.Handler

	Pattern() string
}

func main() {
	fx.New(
		fx.Provide(
			NewHTTPServer,
			NewServeMux,
			fx.Annotate(
				NewEchoHandler,
				fx.As(new(Route)),
			),
			zap.NewExample,
		),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
