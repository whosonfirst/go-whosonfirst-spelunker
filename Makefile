CWD=$(shell pwd)

GOMOD=$(shell test -f "go.work" && echo "readonly" || echo "vendor")
LDFLAGS=-s -w

GOTAGS_SQLITE=sqlite3,icu,json1,fts5
GOTAGS_OPENSEARCH=opensearch

GOTAGS=$(GOTAGS_SQLITE),$(GOTAGS_OPENSEARCH)

cli-sqlite:
	@make cli GOTAGS=$(GOTAGS_SQLITE) 

cli-opensearch:
	@make cli GOTAGS=$(GOTAGS_OPENSEARCH) 

cli:
	go build -mod $(GOMOD) -tags="$(GOTAGS)" -ldflags="$(LDFLAGS)" -o bin/wof-spelunker-httpd cmd/wof-spelunker-httpd/main.go
