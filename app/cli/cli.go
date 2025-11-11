package cli

// To do: model this after the "commands" in whosonfirst/wof-cli

import (
	"context"
	"flag"
	"fmt"

	"github.com/sfomuseum/go-flags/flagset"
	spelunker "github.com/whosonfirst/go-whosonfirst-spelunker/v2"
)

func Run(ctx context.Context) error {
	fs := DefaultFlagSet()
	return RunWithFlagSet(ctx, fs)
}

func RunWithFlagSet(ctx context.Context, fs *flag.FlagSet) error {

	flagset.Parse(fs)

	sp, err := spelunker.NewSpelunker(ctx, spelunker_uri)

	if err != nil {
		return fmt.Errorf("Failed to create new spelunker, %w", err)
	}

	switch command {
	case "descendants":
		return get_descendants(ctx, sp)
	case "id":
		return get_by_id(ctx, sp)
	case "search":
		return search(ctx, sp)
	default:
		return fmt.Errorf("Invalid or unsupported command")
	}

	return nil
}
