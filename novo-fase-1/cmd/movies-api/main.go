package main

import (
	"log"
	"net/http"
	// "novo-fase-1/internal/config"
)

func main() {
	cfg := config.MustLoad()

	store := storage.NewMemoryStorage()

	router := http.NewServeMux()

	// router.HandleFunc()

	log.Printf("servidor rodando em %s (env=%s)", cfg.HTTPServer.Address, cfg.Env)

	if err := http.ListenAndServe(cfg.HTTPServer.Address, router); err != nil {
		log.Fatal(err)
	}

}
