package hasalternateplacetype

import (
	"flag"

	"github.com/sfomuseum/go-flags/flagset"
)

var spelunker_uri string
var pt string
var per_page int64
var page int64

func DefaultFlagSet() *flag.FlagSet {

	fs := flagset.NewFlagSet("getdescendants")
	fs.StringVar(&spelunker_uri, "spelunker-uri", "", "...")
	fs.StringVar(&pt, "placetype", "", "...")
	fs.Int64Var(&page, "page", 1, "...")
	fs.Int64Var(&per_page, "per-page", 10, "...")

	return fs
}
