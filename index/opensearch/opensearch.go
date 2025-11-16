package opensearch

import (
	"context"

	_ "github.com/whosonfirst/go-whosonfirst-database/opensearch/writer"

	"github.com/whosonfirst/go-whosonfirst-iterwriter/v4"
	iterwriter_app "github.com/whosonfirst/go-whosonfirst-iterwriter/v4/app/iterwriter"
	"github.com/whosonfirst/go-whosonfirst-spelunker/v2/index"
)

type IndexOpenSearchCommand struct {
	index.Command
}

func init() {
	ctx := context.Background()
	index.RegisterCommand(ctx, "opensearch", NewIndexOpenSearchCommand)
}

func NewIndexOpenSearchCommand(ctx context.Context, cmd string) (index.Command, error) {
	c := &IndexOpenSearchCommand{}
	return c, nil
}

func (c *IndexOpenSearchCommand) Run(ctx context.Context, args []string) error {

	fs := DefaultFlagSet()
	fs.Parse(args)

	sources := fs.Args()

	cb_func := iterwriter.DefaultIterwriterCallback(forgiving)

	opts := &iterwriter_app.RunOptions{
		CallbackFunc:  cb_func,
		IteratorURI:   iterator_uri,
		IteratorPaths: sources,
		Verbose:       verbose,
	}

	return iterwriter_app.RunWithOptions(ctx, opts)
}
