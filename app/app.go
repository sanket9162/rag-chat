package app

import (
	"context"

	"github.com/sanket9162/rag-chat/config"
	"github.com/sanket9162/rag-chat/llm"
)

func Run(ctx context.Context, cfg config.Config) error {
	client := llm.NewClient(cfg)
	return chat.RunREPL(ctx, client, chat.Options{
		SystemPromptFile: cfg.SystemPromptFile,
	})
}
