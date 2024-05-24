package hasalternateplacetypefaceted

import (
	"context"
	"fmt"
	"os"

	"github.com/whosonfirst/go-whosonfirst-spelunker"
	"github.com/whosonfirst/go-whosonfirst-spelunker/command"
)

type HasAlternatePlacetypeFacetedCommand struct {
	spelunker.Command
}

func init() {
	ctx := context.Background()
	spelunker.RegisterCommand(ctx, "hasalternateplacetypefaceted", NewHasAlternatePlacetypeFacetedCommand)
}

func NewHasAlternatePlacetypeFacetedCommand(ctx context.Context, cmd string) (spelunker.Command, error) {
	c := &HasAlternatePlacetypeFacetedCommand{}
	return c, nil
}

func (c *HasAlternatePlacetypeFacetedCommand) Run(ctx context.Context, args []string) error {

	fs := DefaultFlagSet()
	fs.Parse(args)

	if len(str_facets) == 0 {
		return fmt.Errorf("Missing facets")
	}

	sp, err := spelunker.NewSpelunker(ctx, spelunker_uri)

	if err != nil {
		return fmt.Errorf("Failed to create new spelunker, %w", err)
	}

	filters := make([]spelunker.Filter, 0)

	facets := make([]*spelunker.Facet, len(str_facets))

	for idx, prop := range str_facets {

		f := spelunker.NewFacet(prop)
		facets[idx] = f
	}

	// To do: pagination

	rsp, err := sp.HasAlternatePlacetypeFaceted(ctx, pt, filters, facets)

	if err != nil {
		return fmt.Errorf("Failed to retrieve descendants, %w", err)
	}

	err = command.WriteJSON(rsp, os.Stdout)

	if err != nil {
		return fmt.Errorf("Failed to write JSON, %w", err)
	}

	return nil
}
