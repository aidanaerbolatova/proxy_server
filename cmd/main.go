package main

import (
	"log"
	test_task "proxy"
	cache2 "proxy/internal/cache"
	"proxy/internal/handlers"
	proxy2 "proxy/internal/proxy"
)

const port = "8080"

func main() {
	cache := cache2.NewCache()
	proxy := proxy2.NewProxyServer(cache)
	handler := handlers.NewHandler(proxy)

	log.Printf("Starting server...\nhttp://localhost%v/\n", ":"+port)
	srv := new(test_task.Server)
	if err := srv.Run(port, handler.InitRoutes()); err != nil {
		log.Fatalf("error occured while runnig http server: %s", err.Error())
	}
}
