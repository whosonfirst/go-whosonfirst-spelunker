package opensearch

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	_ "github.com/whosonfirst/go-whosonfirst-database/opensearch/writer"

	"github.com/opensearch-project/opensearch-go/v4/opensearchapi"
	"github.com/whosonfirst/go-whosonfirst-database/opensearch/client"
	"github.com/whosonfirst/go-whosonfirst-database/opensearch/schema/v2"
	"github.com/whosonfirst/go-whosonfirst-iterwriter/v4"
	iterwriter_app "github.com/whosonfirst/go-whosonfirst-iterwriter/v4/app/iterwriter"
	"github.com/whosonfirst/go-whosonfirst-spelunker/v2/index"
	"github.com/whosonfirst/go-writer/v3"
)

type IndexOpenSearchCommand struct {
	index.Command
}

func init() {
	ctx := context.Background()
	index.RegisterCommand(ctx, "opensearch", NewIndexOpenSearchCommand)
}

func NewIndexOpenSearchCommand(ctx context.Context, cmd string) (index.Command, error) {
	c := &IndexOpenSearchCommand{}
	return c, nil
}

func (c *IndexOpenSearchCommand) Run(ctx context.Context, args []string) error {

	fs := DefaultFlagSet()
	fs.Parse(args)

	sources := fs.Args()

	wr, err := writer.NewWriter(ctx, client_uri)

	if err != nil {
		return fmt.Errorf("Failed to create new writer, %w", err)
	}

	u, _ := url.Parse(client_uri)
	os_index := strings.TrimLeft(u.Path, "/")

	if create_index {

		mappings_r, err := v2.FS.Open("mappings.spelunker.json")

		if err != nil {
			return fmt.Errorf("Failed to open mappings for reading, %w", err)
		}

		defer mappings_r.Close()

		settings_r, err := v2.FS.Open("settings.spelunker.json")

		if err != nil {
			return fmt.Errorf("Failed to open settings for reading, %w", err)
		}

		defer settings_r.Close()

		os_client, err := client.NewClient(ctx, client_uri)

		if err != nil {
			return fmt.Errorf("Failed to create Opensearch client, %w", err)
		}

		mappings_req := opensearchapi.IndicesCreateReq{
			Index: os_index,
			Body: mappings_r,
		}
		
		_, err = os_client.Indices.Create(ctx, mappings_req)

		if err != nil {
			return fmt.Errorf("Failed to create index, %w", err)
		}

		settings_req := opensearchapi.SettingsPutReq{
			Indices: []string{
				os_index,
			},
			Body: settings_r,
		}

		_, err = os_client.Indices.Settings.Put(ctx, settings_req)

		if err != nil {
			return fmt.Errorf("Failed to put settings, %w", err)
		}
		
	}

	cb_func := iterwriter.DefaultIterwriterCallback(forgiving)

	opts := &iterwriter_app.RunOptions{
		CallbackFunc:  cb_func,
		Writer:        wr,
		IteratorURI:   iterator_uri,
		IteratorPaths: sources,
		Verbose:       verbose,
	}

	return iterwriter_app.RunWithOptions(ctx, opts)
}
