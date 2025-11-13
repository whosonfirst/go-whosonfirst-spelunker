package www

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"

	"github.com/aaronland/go-http/v4/slog"
	wof_http "github.com/whosonfirst/go-whosonfirst-spelunker/v2/http"
)

type URIsJSHandlerOptions struct {
	Templates *template.Template
	URIs      *wof_http.URIs
}

type URIsJSVars struct {
	Table string
}

func URIsJSHandler(opts *URIsJSHandlerOptions) (http.Handler, error) {

	t := opts.Templates.Lookup("whosonfirst_spelunker_uris")

	if t == nil {
		return nil, fmt.Errorf("Failed to locate 'whosonfirst_spelunker_uris' template")
	}

	fn := func(rsp http.ResponseWriter, req *http.Request) {

		logger := slog.LoggerWithRequest(req, nil)

		enc_table, err := json.Marshal(opts.URIs)

		if err != nil {
			logger.Error("Failed to marshal URIs table", "error", err)
			http.Error(rsp, "Internal server error", http.StatusInternalServerError)
			return
		}

		vars := URIsJSVars{
			Table: string(enc_table),
		}

		rsp.Header().Set("Content-type", "text/javascript")
		err = t.Execute(rsp, vars)

		if err != nil {
			logger.Error("Failed to execute template", "error", err)
			http.Error(rsp, "Internal server error", http.StatusInternalServerError)
			return
		}

		return
	}

	return http.HandlerFunc(fn), nil
}
