package server

import (
	"context"
	"fmt"
	"net/http"

	// "github.com/whosonfirst/go-whosonfirst-spelunker"
	"github.com/whosonfirst/go-whosonfirst-spelunker/http/www"
)

func geoJSONHandlerFunc(ctx context.Context) (http.Handler, error) {

	setupCommonOnce.Do(setupCommon)

	if setupCommonError != nil {
		return nil, fmt.Errorf("Failed to set up common configuration, %w", setupCommonError)
	}

	opts := &www.GeoJSONHandlerOptions{
		Spelunker: sp,
	}

	return www.GeoJSONHandler(opts)
}
