package getbyid

import (
	"context"
	"fmt"

	"github.com/whosonfirst/go-whosonfirst-spelunker"
	"github.com/whosonfirst/go-whosonfirst-uri"
)

type GetByIdCommand struct {
	spelunker.Command
}

func init() {
	ctx := context.Background()
	spelunker.RegisterCommand(ctx, "getbyid", NewGetByIdCommand)
}

func NewGetByIdCommand(ctx context.Context, cmd string) (spelunker.Command, error) {
	c := &GetByIdCommand{}
	return c, nil
}

func (c *GetByIdCommand) Run(ctx context.Context, args []string) error {

	fs := DefaultFlagSet()
	fs.Parse(args)

	sp, err := spelunker.NewSpelunker(ctx, spelunker_uri)

	if err != nil {
		return fmt.Errorf("Failed to create new spelunker, %w", err)
	}

	uri_args := new(uri.URIArgs)

	body, err := sp.GetRecordForId(ctx, id, uri_args)

	if err != nil {
		return fmt.Errorf("Failed to get record by ID, %w", err)
	}

	fmt.Println(string(body))
	return nil
}
