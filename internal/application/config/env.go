package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var (
	once   sync.Once
	config *Config
)

type Config struct {
	ServerAddr string
	ServerPort string
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, loading configuration from environment")
	}

	once.Do(loadConfig)
}

func loadConfig() {
	config = &Config{
		ServerAddr: getenv("ADDRESS", "localhost"),
		ServerPort: getenv("PORT", "4000"),
	}
}

func GetConfig() *Config {
	return config
}

func getenv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
