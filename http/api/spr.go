package api

import (
	"encoding/json"
	"net/http"

	"github.com/aaronland/go-http/v3/slog"
	"github.com/whosonfirst/go-whosonfirst-spelunker"
	wof_http "github.com/whosonfirst/go-whosonfirst-spelunker/http"
	"github.com/whosonfirst/go-whosonfirst-spr/v2"
)

type SPRHandlerOptions struct {
	Spelunker spelunker.Spelunker
}

func SPRHandler(opts *SPRHandlerOptions) (http.Handler, error) {

	fn := func(rsp http.ResponseWriter, req *http.Request) {

		ctx := req.Context()
		logger := slog.LoggerWithRequest(req, nil)

		req_uri, err, status := wof_http.ParseURIFromRequest(req, nil)

		if err != nil {
			logger.Error("Failed to parse URI from request", "error", err)
			http.Error(rsp, spelunker.ErrNotFound.Error(), status)
			return
		}

		if req_uri.Id <= -1 {
			http.Error(rsp, "Not found", http.StatusNotFound)
			return
		}

		logger = logger.With("id", req_uri.Id)

		/*
			spr, err := wof_http.SPRFromRequestURI(ctx, opts.Spelunker, req_uri)

			if err != nil {
				logger.Error("Failed to get by ID", "id", req_uri.Id, "error", err)
				http.Error(rsp, spelunker.ErrNotFound.Error(), http.StatusNotFound)
				return
			}
		*/

		r, err := wof_http.FeatureFromRequestURI(ctx, opts.Spelunker, req_uri)

		if err != nil {
			logger.Error("Failed to get by ID", "error", err)
			http.Error(rsp, spelunker.ErrNotFound.Error(), http.StatusNotFound)
			return
		}

		s, err := spr.WhosOnFirstSPR(r)

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusInternalServerError)
			return
		}

		rsp.Header().Set("Content-Type", "application/json")

		enc := json.NewEncoder(rsp)
		err = enc.Encode(s)

		if err != nil {
			logger.Error("Failed to marshal response", "error", err)
			http.Error(rsp, err.Error(), http.StatusInternalServerError)
			return
		}

		return
	}

	h := http.HandlerFunc(fn)
	return h, nil
}
