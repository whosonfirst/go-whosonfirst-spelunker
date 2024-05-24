package getplacetypes

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/whosonfirst/go-whosonfirst-spelunker"
)

type GetPlacetypesCommand struct {
	spelunker.Command
}

func init() {
	ctx := context.Background()
	spelunker.RegisterCommand(ctx, "getplacetypes", NewGetPlacetypesCommand)
}

func NewGetPlacetypesCommand(ctx context.Context, cmd string) (spelunker.Command, error) {
	c := &GetPlacetypesCommand{}
	return c, nil
}

func (c *GetPlacetypesCommand) Run(ctx context.Context, args []string) error {

	fs := DefaultFlagSet()
	fs.Parse(args)

	sp, err := spelunker.NewSpelunker(ctx, spelunker_uri)

	if err != nil {
		return fmt.Errorf("Failed to create new spelunker, %w", err)
	}

	r, err := sp.GetPlacetypes(ctx)

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
