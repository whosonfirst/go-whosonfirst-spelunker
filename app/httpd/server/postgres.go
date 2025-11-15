//go:build postgres

package server

import (
	_ "github.com/lib/pq"
	_ "github.com/whosonfirst/go-whosonfirst-spelunker/v2/sql"
)
