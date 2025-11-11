package main

import (
	"context"
	"log"

	"github.com/whosonfirst/go-whosonfirst-spelunker/app/cli"
)

func main() {

	ctx := context.Background()
	err := cli.Run(ctx)

	if err != nil {
		log.Fatalf("Failed to run spelunker application, %v", err)
	}
}
