package main

import (
	"fmt"

	"github.com/mcrors/workers/internal/config"
)

func main() {
	cfg, err := config.Load("")
	if err != nil {
		panic(err)
	}
	fmt.Println(cfg)
}
