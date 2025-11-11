CWD=$(shell pwd)

GOMOD=$(shell test -f "go.work" && echo "readonly" || echo "vendor")
LDFLAGS=-s -w

GOTAGS=wof

GOTAGS_SQL=sqlite3,icu,json1,fts5
GOTAGS_OPENSEARCH=opensearch


cli-sql:
	@make cli GOTAGS=$(GOTAGS_SQL) 

cli-opensearch:
	@make cli GOTAGS=$(GOTAGS_OPENSEARCH) 

cli:
	go build -mod $(GOMOD) -tags="$(GOTAGS)" -ldflags="$(LDFLAGS)" -o bin/wof-spelunker cmd/wof-spelunker/main.go
	go build -mod $(GOMOD) -tags="$(GOTAGS)" -ldflags="$(LDFLAGS)" -o bin/wof-spelunker-httpd cmd/wof-spelunker-httpd/main.go
