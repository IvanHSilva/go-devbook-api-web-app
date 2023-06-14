package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

// func init() {
// 	//
// 	key := make([]byte, 64)

// 	if _, err := rand.Read(key); err != nil {
// 		log.Fatal(err)
// 	}

// 	strKey := base64.StdEncoding.EncodeToString(key)
// 	fmt.Println(strKey)
// }

func main() {
	// fmt.Println()
	// fmt.Println("API")
	fmt.Println()

	config.LoadConfig()
	// fmt.Println(config.Connection)
	// fmt.Println(config.Port)
	fmt.Println(config.SecretKey)
	r := router.Generate()

	fmt.Printf("Escutando na porta %d\n", config.Port)
	log.Fatal(http.ListenAndServe(":5000", r))
}
