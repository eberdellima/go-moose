package utilities

import (
	"github.com/joho/godotenv"
	"log"
)

// LoadEnv runs godotenv.Load to load
// environment variables into the process
// Used for testing purposes only
func LoadEnv(envPath string) {

	if err := godotenv.Load(envPath); err != nil {
		log.Fatal(err)
	}
}
