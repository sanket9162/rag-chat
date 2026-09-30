package llm

import (
	"context"

	"github.com/sanket9162/rag-chat/config"
	"google.golang.org/genai"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Client struct {
	config config.Config
	sdk    *genai.Client
}

func NewClient(cfg config.Config) *Client {
	clientConfig := &genai.ClientConfig{
		APIKey: cfg.APIKey,
	}
	if cfg.BaseURL != "" {
		clientConfig.HTTPOptions = genai.HTTPOptions{
			BaseURL: cfg.BaseURL,
		}
	}

	sdk, _ := genai.NewClient(context.Background(), clientConfig)
	return &Client{config: cfg, sdk: sdk}
}
