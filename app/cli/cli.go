package cli

import (
	"context"
	"fmt"
	"os"

	"github.com/whosonfirst/go-whosonfirst-spelunker"
	_ "github.com/whosonfirst/go-whosonfirst-spelunker/command/getalternateplacetypes"
	_ "github.com/whosonfirst/go-whosonfirst-spelunker/command/getbyid"
	_ "github.com/whosonfirst/go-whosonfirst-spelunker/command/getdescendants"
	_ "github.com/whosonfirst/go-whosonfirst-spelunker/command/getplacetypes"
	_ "github.com/whosonfirst/go-whosonfirst-spelunker/command/hasalternateplacetype"
	_ "github.com/whosonfirst/go-whosonfirst-spelunker/command/search"
)

func usage() {

	fmt.Println("Usage: wof-spelunker [CMD] [OPTIONS]")
	fmt.Println("Valid commands are:")

	for _, cmd := range spelunker.Commands() {
		fmt.Printf("* %s\n", cmd)
	}

	os.Exit(0)
}

func Run(ctx context.Context) error {

	if len(os.Args) < 2 {
		usage()
	}

	cmd := os.Args[1]

	c, err := spelunker.NewCommand(ctx, cmd)

	if err != nil {
		usage()
	}

	args := make([]string, 0)

	if len(os.Args) > 2 {
		args = os.Args[2:]
	}

	err = c.Run(ctx, args)

	if err != nil {
		return fmt.Errorf("Failed to run '%s' command, %w", cmd, err)
	}

	return nil
}
