package filters

import (
	"context"

	"github.com/whosonfirst/go-whosonfirst-spelunker"
)

const PLACETYPE_FILTER_SCHEME string = "placetype"

type PlacetypeFilter struct {
	spelunker.Filter
}

func NewPlacetypeFilter(ctx context.Context, uri string) (spelunker.Filter, error) {

	f := &PlacetypeFilter{}
	return f, nil
}

func (f *PlacetypeFilter) Scheme() string {
	return PLACETYPE_FILTER_SCHEME
}

func (f *PlacetypeFilter) Value() any {
	return ""
}
