package handlers

import (
	"net/http"
	"proxy/internal/proxy"
)

type Handler struct {
	proxy *proxy.ProxyServer
}

func NewHandler(proxy *proxy.ProxyServer) *Handler {
	return &Handler{proxy: proxy}
}

func (h *Handler) InitRoutes() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/", h.ProxyServe)

	return router
}
