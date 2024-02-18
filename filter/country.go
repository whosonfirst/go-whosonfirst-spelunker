package filters

import (
	"context"
	"fmt"
	"net/url"

	"github.com/whosonfirst/go-whosonfirst-spelunker"
)

const COUNTRY_FILTER_SCHEME string = "country"

type CountryFilter struct {
	spelunker.Filter
	code string
}

func NewCountryFilter(ctx context.Context, uri string) (spelunker.Filter, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to parse URI, %w", err)
	}

	code := u.Host

	// Validate code here...

	f := &CountryFilter{
		code: code,
	}

	return f, nil
}

func (f *CountryFilter) Scheme() string {
	return COUNTRY_FILTER_SCHEME
}

func (f *CountryFilter) Value() any {
	return f.code
}
