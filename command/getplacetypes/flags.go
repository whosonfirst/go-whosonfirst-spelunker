package getplacetypes

import (
	"flag"

	"github.com/sfomuseum/go-flags/flagset"
)

var spelunker_uri string

func DefaultFlagSet() *flag.FlagSet {

	fs := flagset.NewFlagSet("getplacetypes")
	fs.StringVar(&spelunker_uri, "spelunker-uri", "", "...")

	return fs
}
