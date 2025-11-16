package index

import (
	"context"
	"fmt"
	"os"

	_ "github.com/whosonfirst/go-whosonfirst-spelunker/v2/app"

	"github.com/whosonfirst/go-whosonfirst-spelunker/v2/index"
)

func usage() {

	fmt.Println("Usage: wof-spelunker-index [CMD] [OPTIONS]")
	fmt.Println("Valid commands are:")

	for _, cmd := range index.Commands() {
		fmt.Printf("* %s\n", cmd)
	}

	os.Exit(0)
}

func Run(ctx context.Context) error {

	if len(os.Args) < 2 {
		usage()
	}

	cmd := os.Args[1]

	if cmd == "-h" {
		usage()
	}

	c, err := index.NewCommand(ctx, cmd)

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
