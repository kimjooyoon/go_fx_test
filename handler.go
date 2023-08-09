package main

import (
	"fmt"
	"go.uber.org/zap"
	"io"
	"net/http"
)

type EchoHandler struct {
	log *zap.Logger
}

func (*EchoHandler) Pattern() string {
	return "/echo"
}

func NewEchoHandler(log *zap.Logger) *EchoHandler {
	return &EchoHandler{log}
}

func (h *EchoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, err := io.Copy(w, r.Body); err != nil {
		h.log.Warn("Failed to handle request", zap.Error(err))
	}
}

func NewServeMux(route1, route2 Route) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle(route1.Pattern(), route1)
	mux.Handle(route2.Pattern(), route2)
	return mux
}

type HelloHandler struct {
	log *zap.Logger
}

func NewHelloHandler(log *zap.Logger) *HelloHandler {
	return &HelloHandler{log: log}
}
func (*HelloHandler) Pattern() string {
	return "/hello"
}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.log.Error("Failed to read request", zap.Error(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if _, err := fmt.Fprintf(w, "Hello, %s\n", body); err != nil {
		h.log.Error("Failed to write response", zap.Error(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
