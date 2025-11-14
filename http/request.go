package http

import (
	"context"
	"fmt"
	go_http "net/http"

	"github.com/aaronland/go-http/v4/sanitize"
	"github.com/aaronland/go-pagination"
	"github.com/aaronland/go-pagination/countable"
	"github.com/aaronland/go-pagination/cursor"
	"github.com/whosonfirst/go-whosonfirst-spelunker/v2"
	"github.com/whosonfirst/go-whosonfirst-spr/v2"
	wof_http "github.com/whosonfirst/go-whosonfirst/http"
)

func PaginationOptionsFromRequest(req *go_http.Request) (pagination.Options, error) {

	q_cursor, err := sanitize.GetString(req, "cursor")

	if err != nil {
		return nil, fmt.Errorf("Failed to derive ?cursor= parameter, %w", err)
	}

	if q_cursor != "" {

		pg_opts, err := cursor.NewCursorOptions()

		if err != nil {
			return nil, fmt.Errorf("Failed to create cursor options, %w", err)
		}

		pg_opts.Pointer(q_cursor)
		return pg_opts, nil
	}

	page, err := sanitize.GetInt64(req, "page")

	if err != nil {
		return nil, fmt.Errorf("Failed to derive ?page= parameter, %w", err)
	}

	if page == 0 {
		page = 1
	}

	pg_opts, err := countable.NewCountableOptions()

	if err != nil {
		return nil, fmt.Errorf("Failed to create countable options, %w", err)
	}

	pg_opts.Pointer(page)
	return pg_opts, nil
}

func ParsePageNumberFromRequest(req *go_http.Request) (int64, error) {

	page, err := sanitize.GetInt64(req, "page")

	if err != nil {
		return 0, fmt.Errorf("Failed to derive ?page= parameter, %w", err)
	}

	if page == 0 {
		page = 1
	}

	return page, nil
}

func FeatureFromRequestURI(ctx context.Context, sp spelunker.Spelunker, req_uri *wof_http.URI) ([]byte, error) {

	wof_id := req_uri.Id

	f, err := sp.GetFeatureForId(ctx, wof_id, req_uri.URIArgs)

	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve feature for %d, %w", wof_id, err)
	}

	return f, nil
}

func RecordFromRequestURI(ctx context.Context, sp spelunker.Spelunker, req_uri *wof_http.URI) ([]byte, error) {

	wof_id := req_uri.Id

	f, err := sp.GetRecordForId(ctx, wof_id, req_uri.URIArgs)

	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve record for %d, %w", wof_id, err)
	}

	return f, nil
}

func SPRFromRequestURI(ctx context.Context, sp spelunker.Spelunker, req_uri *wof_http.URI) (spr.StandardPlacesResult, error) {

	wof_id := req_uri.Id

	f, err := sp.GetSPRForId(ctx, wof_id, req_uri.URIArgs)

	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve record for %d, %w", wof_id, err)
	}

	return f, nil
}
