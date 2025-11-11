//go:build sqlite3

package server

import (
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/whosonfirst/go-whosonfirst-spelunker/sql"
)
