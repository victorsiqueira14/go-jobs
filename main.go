package main // declaração do pacote principal

import (
	"github.com/victorsiqueira14/go-jobs/config"
	"github.com/victorsiqueira14/go-jobs/router"
)

var (
	logger *config.Logger
)

// arquivo que contém a função main e inicializa tudo
func main() {
	logger = config.GetLogger("main")
	err := config.Init()

	// initialize config
	if err != nil {
		logger.Errorf("Config initializing error: %v", err)
		return
	}

	// initialize router
	router.Initialize()
}
