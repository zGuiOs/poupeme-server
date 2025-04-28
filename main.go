package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/zGuiOs/poupeme-server/src/config"
	"github.com/zGuiOs/poupeme-server/src/router"
)

func main() {
	config.Load()

	fmt.Println("Rodando")

	router := router.Build()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.ApiPort), router))
}
