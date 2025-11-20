package www

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/aaronland/go-http/v4/auth"
	"github.com/aaronland/go-http/v4/slog"
	wof_http "github.com/whosonfirst/go-whosonfirst-spelunker/v2/http"
)

type TemplateHandlerOptions struct {
	Authenticator auth.Authenticator
	Templates     *template.Template
	TemplateName  string
	PageTitle     string
	URIs          *wof_http.URIs
}

type templateHandlerVars struct {
	Id         int64
	PageTitle  string
	URIs       *wof_http.URIs
	Properties string
	OpenGraph  *OpenGraph
}

func TemplateHandler(opts *TemplateHandlerOptions) (http.Handler, error) {

	t := opts.Templates.Lookup(opts.TemplateName)

	if t == nil {
		return nil, fmt.Errorf("Failed to locate ihelp' template")
	}

	fn := func(rsp http.ResponseWriter, req *http.Request) {

		logger := slog.LoggerWithRequest(req, nil)

		vars := templateHandlerVars{
			PageTitle: opts.PageTitle,
			URIs:      opts.URIs,
		}

		vars.OpenGraph = &OpenGraph{
			Type:        "Article",
			SiteName:    "Who's On First Spelunker",
			Title:       fmt.Sprintf("Who's On First Spelunker – %s", opts.PageTitle),
			Description: "",
			Image:       "",
		}

		rsp.Header().Set("Content-Type", "text/html")

		err := t.Execute(rsp, vars)

		if err != nil {
			logger.Error("Failed to render template ", "error", err)
			http.Error(rsp, "Internal server error", http.StatusInternalServerError)
		}

	}

	h := http.HandlerFunc(fn)
	return h, nil
}
