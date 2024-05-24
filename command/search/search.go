package search

import (
	"context"
	"encoding/json"
	"fmt"
	_ "log/slog"
	"os"

	"github.com/aaronland/go-pagination/countable"
	"github.com/whosonfirst/go-whosonfirst-spelunker"
)

type SearchCommand struct {
	spelunker.Command
}

func init() {
	ctx := context.Background()
	spelunker.RegisterCommand(ctx, "search", NewSearchCommand)
}

func NewSearchCommand(ctx context.Context, cmd string) (spelunker.Command, error) {
	c := &SearchCommand{}
	return c, nil
}

func (c *SearchCommand) Run(ctx context.Context, args []string) error {

	fs := DefaultFlagSet()
	fs.Parse(args)

	sp, err := spelunker.NewSpelunker(ctx, spelunker_uri)

	if err != nil {
		return fmt.Errorf("Failed to create new spelunker, %w", err)
	}

	// Eventually we'll need to check if we're doing cursor-base pagination

	pg_opts, err := countable.NewCountableOptions()

	if err != nil {
		return fmt.Errorf("Failed to create countable options, %w", err)
	}

	pg_opts.PerPage(per_page)
	pg_opts.Pointer(page)

	search_opts := &spelunker.SearchOptions{
		Query: query,
	}

	filters := make([]spelunker.Filter, 0)

	r, _, err := sp.Search(ctx, pg_opts, search_opts, filters)

	if err != nil {
		return fmt.Errorf("Failed to retrieve descendants, %w", err)
	}

	enc := json.NewEncoder(os.Stdout)
	err = enc.Encode(r)

	if err != nil {
		return fmt.Errorf("Failed to encode descendants, %w", err)
	}

	return nil
}
