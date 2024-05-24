package hasalternateplacetype

import (
	"context"
	"fmt"
	"os"

	"github.com/aaronland/go-pagination/countable"
	"github.com/whosonfirst/go-whosonfirst-spelunker"
	"github.com/whosonfirst/go-whosonfirst-spelunker/command"
)

type HasAlternatePlacetypeCommand struct {
	spelunker.Command
}

func init() {
	ctx := context.Background()
	spelunker.RegisterCommand(ctx, "hasalternateplacetype", NewHasAlternatePlacetypeCommand)
}

func NewHasAlternatePlacetypeCommand(ctx context.Context, cmd string) (spelunker.Command, error) {
	c := &HasAlternatePlacetypeCommand{}
	return c, nil
}

func (c *HasAlternatePlacetypeCommand) Run(ctx context.Context, args []string) error {

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

	filters := make([]spelunker.Filter, 0)

	// To do: pagination

	rsp, _, err := sp.HasAlternatePlacetype(ctx, pg_opts, pt, filters)

	if err != nil {
		return fmt.Errorf("Failed to retrieve descendants, %w", err)
	}

	command.WriteSPRResults(rsp, os.Stdout)
	return nil
}
