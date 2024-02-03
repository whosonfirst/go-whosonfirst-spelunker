package main

import (
	"context"
	"log/slog"
	"os"

	_ "github.com/whosonfirst/go-reader-findingaid"
	"github.com/whosonfirst/go-whosonfirst-spelunker/app/server"
)

func main() {

	ctx := context.Background()
	logger := slog.Default()

	err := server.Run(ctx, logger)

	if err != nil {
		slog.Error("Failed to run server", "error", err)
		os.Exit(1)
	}
}
