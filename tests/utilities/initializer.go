package utilities

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnv(envPath string) {

	if err := godotenv.Load(envPath); err != nil {
		log.Fatal(err)
	}
}
