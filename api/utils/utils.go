package utils

import (
	"github.com/joho/godotenv"
)

// LoadEnv will load the environment variables
func LoadEnv() {
	godotenv.Load()
}
