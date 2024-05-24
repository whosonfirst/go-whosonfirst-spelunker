package search

import (
	"flag"

	"github.com/sfomuseum/go-flags/flagset"
)

var spelunker_uri string
var per_page int64
var page int64
var query string

func DefaultFlagSet() *flag.FlagSet {

	fs := flagset.NewFlagSet("search")
	fs.StringVar(&spelunker_uri, "spelunker-uri", "", "...")
	fs.Int64Var(&page, "page", 1, "...")
	fs.Int64Var(&per_page, "per-page", 10, "...")
	fs.StringVar(&query, "query", "", "...")
	return fs
}
