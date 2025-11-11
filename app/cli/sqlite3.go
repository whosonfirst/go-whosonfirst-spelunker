//go:build sqlite3

package cli

import (
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/whosonfirst/go-whosonfirst-spelunker/sql"
)
