package spelunker

import (
	"context"
	"fmt"
	"net/url"
)

const PLACETYPE_ALT_FILTER_SCHEME string = "placetype"

type PlacetypeAltFilter struct {
	Filter
	placetype string
}

func NewPlacetypeAltFilterFromString(ctx context.Context, name string) (Filter, error) {
	uri := fmt.Sprintf("%s://%s", PLACETYPE_ALT_FILTER_SCHEME, name)
	return NewPlacetypeAltFilter(ctx, uri)
}

func NewPlacetypeAltFilter(ctx context.Context, uri string) (Filter, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to parse URI, %w", err)
	}

	pt := u.Host

	f := &PlacetypeAltFilter{
		placetype: pt,
	}

	return f, nil
}

func (f *PlacetypeAltFilter) Scheme() string {
	return PLACETYPE_ALT_FILTER_SCHEME
}

func (f *PlacetypeAltFilter) Value() any {
	return f.placetype
}
