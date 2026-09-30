package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	BaseURL          string
	Model            string
	APIKey           string
	SystemPromptFile string
}

func Load() Config {

	_ = godotenv.Load()

	cfg := Config{
		BaseURL:          os.Getenv("BASE_URL"),
		Model:            os.Getenv("MODEL"),
		APIKey:           os.Getenv("API_KEY"),
		SystemPromptFile: os.Getenv("SYSTEM_PROMPT_FILE"),
	}

	if cfg.BaseURL == "" || cfg.APIKey == "" || cfg.Model == "" || cfg.SystemPromptFile == "" {
		log.Fatalf("missing required environment variables")
	}

	return cfg
}
