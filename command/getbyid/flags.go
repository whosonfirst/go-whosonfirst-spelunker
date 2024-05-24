package getbyid

import (
	"flag"

	"github.com/sfomuseum/go-flags/flagset"
)

var spelunker_uri string
var id int64

func DefaultFlagSet() *flag.FlagSet {

	fs := flagset.NewFlagSet("getbyid")
	fs.StringVar(&spelunker_uri, "spelunker-uri", "", "...")
	fs.Int64Var(&id, "id", 0, "...")

	return fs
}
