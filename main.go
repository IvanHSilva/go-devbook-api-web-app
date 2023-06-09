package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// fmt.Println()
	// fmt.Println("API")
	fmt.Println()

	config.LoadConfig()
	// fmt.Println(config.Connection)
	// fmt.Println(config.Port)
	r := router.Generate()

	fmt.Printf("Escutando na porta %d\n", config.Port)
	log.Fatal(http.ListenAndServe(":5000", r))
}
