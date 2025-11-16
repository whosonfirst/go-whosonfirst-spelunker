package sql

import (
	"context"

	sql_index "github.com/whosonfirst/go-whosonfirst-database/app/sql/tables/index"
	"github.com/whosonfirst/go-whosonfirst-spelunker/v2/index"
)

type IndexSQLCommand struct {
	index.Command
}

func init() {
	ctx := context.Background()
	index.RegisterCommand(ctx, "sql", NewIndexSQLCommand)
}

func NewIndexSQLCommand(ctx context.Context, cmd string) (index.Command, error) {
	c := &IndexSQLCommand{}
	return c, nil
}

func (c *IndexSQLCommand) Run(ctx context.Context, args []string) error {

	fs := DefaultFlagSet()
	fs.Parse(args)

	sources := fs.Args()

	opts := &sql_index.RunOptions{
		SpelunkerTables: true,
		DatabaseURI:     db_uri,
		IteratorURI:     iterator_uri,
		IteratorSources: sources,
		MaxProcesses:    procs,
		Verbose:         verbose,
	}

	return sql_index.RunWithOptions(ctx, opts)
}
