package index

import (
	"context"
	"log/slog"
	// sql_index "github.com/whosonfirst/go-whosonfirst-database/app/sql/tables/index"
)

type IndexSQLCommand struct {
	Command
}

func init() {
	ctx := context.Background()
	RegisterCommand(ctx, "sql", NewIndexSQLCommand)
}

func NewIndexSQLCommand(ctx context.Context, cmd string) (Command, error) {
	c := &IndexSQLCommand{}
	return c, nil
}

func (c *IndexSQLCommand) Run(ctx context.Context, args []string) error {

	slog.Info("WOO")
	return nil
}
