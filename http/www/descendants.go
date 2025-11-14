package www

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/aaronland/go-http/v4/auth"
	"github.com/aaronland/go-http/v4/slog"
	"github.com/aaronland/go-pagination"
	"github.com/aaronland/go-pagination/countable"
	"github.com/whosonfirst/go-whosonfirst-spelunker/v2"
	"github.com/whosonfirst/go-whosonfirst-spr/v2"
	wof_http "github.com/whosonfirst/go-whosonfirst/http"
)

type DescendantsHandlerOptions struct {
	Spelunker     spelunker.Spelunker
	Authenticator auth.Authenticator
	Templates     *template.Template
	URIs          *wof_http.URIs
}

type DescendantsHandlerVars struct {
	PageTitle        string
	Id               int64
	URIs             *wof_http.URIs
	Places           []spr.StandardPlacesResult
	Pagination       pagination.Results
	PaginationURL    string
	FacetsURL        string
	FacetsContextURL string
}

func DescendantsHandler(opts *DescendantsHandlerOptions) (http.Handler, error) {

	t := opts.Templates.Lookup("descendants")

	if t == nil {
		return nil, fmt.Errorf("Failed to locate 'descendants' template")
	}

	fn := func(rsp http.ResponseWriter, req *http.Request) {

		ctx := req.Context()
		logger := slog.LoggerWithRequest(req, nil)

		uri, err, status := wof_http.ParseURIFromRequest(req, nil)

		if err != nil {
			logger.Error("Failed to parse URI from request", "error", err)
			http.Error(rsp, spelunker.ErrNotFound.Error(), status)
			return
		}

		logger = logger.With("wofid", uri.Id)

		pg_opts, err := countable.NewCountableOptions()

		if err != nil {
			logger.Error("Failed to create pagination options", "error", err)
			http.Error(rsp, "womp womp", http.StatusInternalServerError)
			return
		}

		pg, pg_err := wof_http.ParsePageNumberFromRequest(req)

		if pg_err == nil {
			pg_opts.Pointer(pg)
		}

		filter_params := wof_http.DefaultFilterParams()

		filters, err := wof_http.FiltersFromRequest(ctx, req, filter_params)

		if err != nil {
			logger.Error("Failed to derive filters from request", "error", err)
			http.Error(rsp, "Bad request", http.StatusBadRequest)
			return
		}

		r, pg_r, err := opts.Spelunker.GetDescendants(ctx, pg_opts, uri.Id, filters)

		if err != nil {
			logger.Error("Failed to get descendants", "error", err)
			http.Error(rsp, "womp womp", http.StatusInternalServerError)
			return
		}

		// This is not ideal but I am not sure what is better yet...
		pagination_url := wof_http.URIForId(opts.URIs.Descendants, uri.Id, filters, nil)

		// This is not ideal but I am not sure what is better yet...
		facets_url := wof_http.URIForId(opts.URIs.DescendantsFaceted, uri.Id, filters, nil)
		facets_context_url := pagination_url

		vars := DescendantsHandlerVars{
			Id:               uri.Id,
			Places:           r.Results(),
			Pagination:       pg_r,
			URIs:             opts.URIs,
			PaginationURL:    pagination_url,
			FacetsURL:        facets_url,
			FacetsContextURL: facets_context_url,
		}

		rsp.Header().Set("Content-Type", "text/html")

		err = t.Execute(rsp, vars)

		if err != nil {
			logger.Error("Failed to return ", "error", err)
			http.Error(rsp, "womp womp", http.StatusInternalServerError)
		}

	}

	h := http.HandlerFunc(fn)
	return h, nil
}
