package api

import (
	"net/http"

	"github.com/aaronland/go-http/v4/slog"
	"github.com/whosonfirst/go-whosonfirst-spelunker/v2"
	wof_http "github.com/whosonfirst/go-whosonfirst-spelunker/v2/http"
)

type GeoJSONHandlerOptions struct {
	Spelunker spelunker.Spelunker
}

func GeoJSONHandler(opts *GeoJSONHandlerOptions) (http.Handler, error) {

	fn := func(rsp http.ResponseWriter, req *http.Request) {

		ctx := req.Context()
		logger := slog.LoggerWithRequest(req, nil)

		req_uri, err, status := wof_http.ParseURIFromRequest(req, nil)

		if err != nil {
			logger.Error("Failed to parse URI from request", "error", err)
			http.Error(rsp, spelunker.ErrNotFound.Error(), status)
			return
		}

		wof_id := req_uri.Id

		if wof_id <= -1 {
			http.Error(rsp, "Not found", http.StatusNotFound)
			return
		}

		logger = logger.With("wof id", wof_id)

		r, err := wof_http.FeatureFromRequestURI(ctx, opts.Spelunker, req_uri)

		if err != nil {
			logger.Error("Failed to get by ID", "error", err)
			http.Error(rsp, spelunker.ErrNotFound.Error(), http.StatusNotFound)
			return
		}

		rsp.Header().Set("Content-Type", "application/json")
		rsp.Write(r)
	}

	h := http.HandlerFunc(fn)
	return h, nil
}
