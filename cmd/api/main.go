package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/yjhss/opsentinel/internal/api"
)

func main() {
	addr := os.Getenv("HTTP_ADDR")
	if addr == "" {
		addr = "127.0.0.1:8080"
	}

	version := os.Getenv("APP_VERSION")
	if version == "" {
		version = "dev"
	}

	server := &http.Server{
		Addr:              addr,
		Handler:           api.NewHandler(version),
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("Opsentinel API 시작: addr=%s version=%s", addr, version)

	if err := server.ListenAndServe(); err != nil &&
		!errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}
