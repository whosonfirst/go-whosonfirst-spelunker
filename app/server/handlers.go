package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
)

func idHandlerFunc(ctx context.Context) (http.Handler, error) {

	setupCommonOnce.Do(setupCommon)

	if setupCommonError != nil {
		return nil, fmt.Errorf("Failed to set up common configuration, %w", setupCommonError)
	}

	fn := func(rsp http.ResponseWriter, req *http.Request) {
		slog.Info("HELLO WORLD")
		return
	}

	return http.HandlerFunc(fn), nil
}
