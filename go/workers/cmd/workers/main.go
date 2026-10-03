package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mcrors/workers/internal/config"
)

func main() {
	cfg, err := config.Load("")
	if err != nil {
		log.Fatalf("config: %v", err)
	}

	mux := http.NewServeMux()
	if err :=
}
