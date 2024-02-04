package cli

import (
	"context"
	"log/slog"

	spelunker "github.com/whosonfirst/go-whosonfirst-spelunker"
)

func get_descendants(ctx context.Context, sp spelunker.Spelunker) error {

	_, _, err := sp.GetDescendants(ctx, id, nil)

	slog.Info("D", "error", err)
	return nil
}
