package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL    string
	ContextTimeout int
}

func LoadConfig() (config *Config) {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	databaseURL := os.Getenv("DATABASE_URL")
	contextTimeout, _ := strconv.Atoi(os.Getenv("CONTEXT_TIMEOUT"))

	return &Config{
		DatabaseURL:    databaseURL,
		ContextTimeout: contextTimeout,
	}
}
