package server

import (
	html_template "html/template"
	"sync"

	"github.com/aaronland/go-http/v3/auth"
	"github.com/rs/cors"
	"github.com/whosonfirst/go-whosonfirst-spelunker"
	wof_http "github.com/whosonfirst/go-whosonfirst-spelunker/http"
)

var run_options *RunOptions

var sp spelunker.Spelunker

var authenticator auth.Authenticator

var uris_table *wof_http.URIs

var html_templates *html_template.Template

var setupCommonOnce sync.Once
var setupCommonError error

var setupWWWOnce sync.Once
var setupWWWError error

var setupAPIOnce sync.Once
var setupAPIError error

var cors_wrapper *cors.Cors
