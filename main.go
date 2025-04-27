package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/zGuiOs/poupeme-server/src/router"
)

func main() {
	fmt.Println("Rodando")

	router := router.Build()

	log.Fatal(http.ListenAndServe(":5000", router))
}
