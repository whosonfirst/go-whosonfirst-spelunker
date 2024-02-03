package spelunker

import (
	"context"
	"fmt"
	_ "log/slog"
	"net/url"

	"github.com/whosonfirst/go-reader"
	wof_reader "github.com/whosonfirst/go-whosonfirst-reader"
)

type ReaderSpelunker struct {
	Spelunker
	reader reader.Reader
}

func init() {
	ctx := context.Background()
	RegisterSpelunker(ctx, "reader", NewReaderSpelunker)
}

func NewReaderSpelunker(ctx context.Context, uri string) (Spelunker, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to parse URI, %w", err)
	}

	q := u.Query()

	r, err := reader.NewReader(ctx, q.Get("reader-uri"))

	if err != nil {
		return nil, fmt.Errorf("Failed to create new reader, %w", err)
	}

	s := &ReaderSpelunker{
		reader: r,
	}

	return s, nil
}

func (s *ReaderSpelunker) GetById(ctx context.Context, id int64) ([]byte, error) {

	r, err := wof_reader.LoadBytes(ctx, s.reader, id)

	if err != nil {
		return nil, fmt.Errorf("Failed to read %d, %w", id, err)
	}

	return r, nil
}

func (s *ReaderSpelunker) GetDescendants(ctx context.Context, id int64) ([][]byte, error) {
	return nil, ErrNotImplemented
}
