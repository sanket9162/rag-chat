package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/sanket9162/rag-chat/app"
	"github.com/sanket9162/rag-chat/config"
)

func main() {

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	if err := app.Run(ctx, config.Load()); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}
