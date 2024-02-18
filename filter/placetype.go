package filters

import (
	"context"
	"fmt"
	"net/url"

	"github.com/whosonfirst/go-whosonfirst-placetypes"
	"github.com/whosonfirst/go-whosonfirst-spelunker"
)

const PLACETYPE_FILTER_SCHEME string = "placetype"

type PlacetypeFilter struct {
	spelunker.Filter
	placetype string
}

func NewPlacetypeFilter(ctx context.Context, uri string) (spelunker.Filter, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to parse URI, %w", err)
	}

	pt := u.Host

	if !placetypes.IsValidPlacetype(pt) {
		return nil, fmt.Errorf("Invalid placetype")
	}

	f := &PlacetypeFilter{
		placetype: pt,
	}

	return f, nil
}

func (f *PlacetypeFilter) Scheme() string {
	return PLACETYPE_FILTER_SCHEME
}

func (f *PlacetypeFilter) Value() any {
	return f.placetype
}
