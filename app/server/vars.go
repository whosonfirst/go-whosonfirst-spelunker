package server

import (
	"sync"

	"github.com/whosonfirst/go-whosonfirst-spelunker"
	sp_http "github.com/whosonfirst/go-whosonfirst-spelunker/http"
)

var sp spelunker.Spelunker

var uris_table *sp_http.URIs

var setupCommonOnce sync.Once
var setupCommonError error
