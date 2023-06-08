package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	Connection = ""
	Port       = 0
)

func LoadConfig() {
	var err error

	if godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Port = 0
	}

	//Connection = "'Server=SERVIDOR\\SQLSERVER;Database=eCommerce;Trusted Connection=true;'"
	fmt.Println(os.Getenv("DB_SERVER"))
	fmt.Println(os.Getenv("DB_NAME"))
	Connection = fmt.Sprintf("'Server=%s;Database=%s;Trusted Connection=true'", os.Getenv("DB_SERVER"), os.Getenv("DB_NAME"))
}
