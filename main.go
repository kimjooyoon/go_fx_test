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

type Params struct {
	fx.In

	Logger *zap.Logger `optional:"true"`
}

func main() {
	fx.New(
		fx.Provide(
			NewHTTPServer,
			fx.Annotate(
				NewServeMux,
				fx.ParamTags(`group:"routes"`),
			),
			AsRoute(NewEchoHandler),
			AsRoute(NewHelloHandler),
		),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}

func AsRoute(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(Route)),
		fx.ResultTags(`group:"routes"`),
	)
}
