package getalternateplacetypes

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/whosonfirst/go-whosonfirst-spelunker"
)

type GetAlternatePlacetypesCommand struct {
	spelunker.Command
}

func init() {
	ctx := context.Background()
	spelunker.RegisterCommand(ctx, "getalternateplacetypes", NewGetAlternatePlacetypesCommand)
}

func NewGetAlternatePlacetypesCommand(ctx context.Context, cmd string) (spelunker.Command, error) {
	c := &GetAlternatePlacetypesCommand{}
	return c, nil
}

func (c *GetAlternatePlacetypesCommand) Run(ctx context.Context, args []string) error {

	fs := DefaultFlagSet()
	fs.Parse(args)

	sp, err := spelunker.NewSpelunker(ctx, spelunker_uri)

	if err != nil {
		return fmt.Errorf("Failed to create new spelunker, %w", err)
	}

	r, err := sp.GetAlternatePlacetypes(ctx)

	if err != nil {
		return fmt.Errorf("Failed to retrieve placetypes, %w", err)
	}

	enc := json.NewEncoder(os.Stdout)
	err = enc.Encode(r)

	if err != nil {
		return fmt.Errorf("Failed to encode placetypes, %w", err)
	}

	return nil
}
