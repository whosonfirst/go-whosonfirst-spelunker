package api

import (
	"net/http"

	"github.com/aaronland/go-http/v3/slog"
	"github.com/whosonfirst/go-whosonfirst-spelunker"
	"github.com/whosonfirst/go-whosonfirst-spelunker/http"
)

type FindingAidHandlerOptions struct {
	Spelunker spelunker.Spelunker
}

func FindingAidHandler(opts *FindingAidHandlerOptions) (http.Handler, error) {

	fn := func(rsp http.ResponseWriter, req *http.Request) {

		ctx := req.Context()
		logger := slog.LoggerWithRequest(req, nil)

		req_uri, err, status := http.ParseURIFromRequest(req, nil)

		if err != nil {
			logger.Error("Failed to parse URI from request", "error", err)
			http.Error(rsp, spelunker.ErrNotFound.Error(), status)
			return
		}

		spr, err := http.SPRFromRequestURI(ctx, opts.Spelunker, req_uri)

		if err != nil {
			logger.Error("Failed to get by ID", "id", req_uri.Id, "error", err)
			http.Error(rsp, spelunker.ErrNotFound.Error(), http.StatusNotFound)
			return
		}

		repo := spr.Repo()

		rsp.Header().Set("Content-Type", "text/plain")
		rsp.Write([]byte(repo))
	}

	h := http.HandlerFunc(fn)
	return h, nil
}
