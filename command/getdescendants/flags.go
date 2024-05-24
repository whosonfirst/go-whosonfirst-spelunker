package getdescendants

import (
	"flag"

	"github.com/sfomuseum/go-flags/flagset"
)

var spelunker_uri string
var id int64
var per_page int64
var page int64

func DefaultFlagSet() *flag.FlagSet {

	fs := flagset.NewFlagSet("getdescendants")
	fs.StringVar(&spelunker_uri, "spelunker-uri", "", "...")
	fs.Int64Var(&id, "id", 0, "...")
	fs.Int64Var(&page, "page", 1, "...")
	fs.Int64Var(&per_page, "per-page", 10, "...")

	return fs
}
