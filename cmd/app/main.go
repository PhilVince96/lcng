package main

import (
	"fmt"
	"lcng/app/router"
	"log"
	"net/http"

	"lcng/config"
)

func main() {
	appConf := config.AppConfig()

	appRouter := router.New()

	address := fmt.Sprintf(":%d", appConf.Server.Port)

	log.Println("Starting app :8080")

	s := &http.Server{
		Addr:         address,
		Handler:      appRouter,
		ReadTimeout:  appConf.Server.TimeoutRead,
		WriteTimeout: appConf.Server.TimeoutWrite,
		IdleTimeout:  appConf.Server.TimeoutIdle,
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server startup failed")
	}
}

func Greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
