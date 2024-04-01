package main

import (
	"github.com/zGuiOs/poupeme-server/config"
	"github.com/zGuiOs/poupeme-server/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	// Iniciando configs
	err := config.Init()

	if err != nil {
		logger.Errorf("Erro ao iniciar as configs, error: %v", err)
		return
	}

	// Iniciando router
	router.Initialize()
}
