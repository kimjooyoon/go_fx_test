package main

import (
	"go.uber.org/fx"
	"net/http"
)

func main() {
	fx.New(
		fx.Provide(
			NewHTTPServer,
			NewServeMux,
			NewEchoHandler,
		),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
