package spelunker

import (
	"context"
	"testing"
)

func TestPlacetypeFilter(t *testing.T) {

	ctx := context.Background()

	ok := []string{
		"placetype://country",
		"placetype://region",
		"placetype://microhood",
		"placetype://custom",
	}

	not_ok := []string{
		"placetype://airport",
	}

	for _, uri := range ok {

		_, err := NewPlacetypeFilter(ctx, uri)

		if err != nil {
			t.Fatalf("Failed to create new filter for %s, %v", uri, err)
		}
	}

	for _, uri := range not_ok {

		_, err := NewPlacetypeFilter(ctx, uri)

		if err == nil {
			t.Fatalf("Did not expect to create new filter for %s", uri)
		}
	}

}
