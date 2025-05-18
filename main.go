package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {

		log.Fatal("Arquivo não encontrado .env !")
	}

	fmt.Println(os.Getenv("test"))

}
