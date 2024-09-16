package main

import (
	"fmt"
	"log"
	"net/http"

	"go-xsd/api"
	"go-xsd/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Erro ao carregar a configuração: %v", err)
	}

	r := api.NewRouter(cfg)

	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	log.Printf("Iniciando servidor na porta %s...", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
