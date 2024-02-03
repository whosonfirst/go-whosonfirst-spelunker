package www

import (
	"log/slog"
	"net/http"

	"github.com/whosonfirst/go-whosonfirst-spelunker"
	sp_http "github.com/whosonfirst/go-whosonfirst-spelunker/http"
)

type GeoJSONHandlerOptions struct {
	Spelunker spelunker.Spelunker
}

func GeoJSONHandler(opts *GeoJSONHandlerOptions) (http.Handler, error) {

	fn := func(rsp http.ResponseWriter, req *http.Request) {

		ctx := req.Context()

		logger := slog.Default()
		logger = logger.With("request", req.URL)

		uri, err, status := sp_http.ParseURIFromRequest(req, nil)

		if err != nil {
			slog.Error("Failed to parse URI from request", "error", err)
			http.Error(rsp, spelunker.ErrNotFound.Error(), status)
			return
		}

		r, err := opts.Spelunker.GetById(ctx, uri.Id)

		if err != nil {
			slog.Error("Failed to get by ID", "id", uri.Id)
			http.Error(rsp, spelunker.ErrNotFound.Error(), http.StatusNotFound)
			return
		}

		rsp.Header().Set("Content-Type", "application/json")
		rsp.Write(r)
	}

	h := http.HandlerFunc(fn)
	return h, nil
}
