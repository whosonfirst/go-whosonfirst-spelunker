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
	go build -mod $(GOMOD) -tags="$(GOTAGS)" -ldflags="$(LDFLAGS)" -o bin/wof-spelunker-index cmd/wof-spelunker-index/main.go

# Targets for running the Spelunker locally

# https://github.com/whosonfirst/go-whosonfirst-database
OS_INDEX=/usr/local/whosonfirst/go-whosonfirst-database/bin/wof-opensearch-index

# https://github.com/whosonfirst/whosonfirst-database
WHOSONFIRST_OPENSEARCH=/usr/local/whosonfirst/go-whosonfirst-database/opensearch

# https://github.com/aaronland/go-tools
URLESCAPE=$(shell which urlescape)

# Opensearch server

# This is for debugging. Do not change this at your own risk.
# (That means you should change this.)
OS_PSWD=dkjfhsjdkfkjdjhksfhskd98475kjHkzjxckj

OS_CACHE_URI=ristretto://
OS_ENC_CACHE_URI=$(shell $(URLESCAPE) $(OS_CACHE_URI))

OS_READER_URI=https://data.whosonfirst.org
OS_ENC_READER_URI=$(shell $(URLESCAPE) $(OS_READER_URI))

OS_CLIENT_URI="https://localhost:9200/spelunker?username=admin&password=$(OS_PSWD)&insecure=true&require-tls=true"
OS_ENC_CLIENT_URI=$(shell $(URLESCAPE) $(OS_CLIENT_URI))

OS_SPELUNKER_URI=opensearch://?client-uri=$(OS_ENC_CLIENT_URI)&cache-uri=$(OS_ENC_CACHE_URI)&reader-uri=$(OS_ENC_READER_URI)

# https://opensearch.org/docs/latest/install-and-configure/install-opensearch/docker/
#
# And then:
# curl -v -k https://admin:$(OS_PSWD)@localhost:9200/

os-local:
	docker run \
		-it \
		-p 9200:9200 \
		-p 9600:9600 \
		-e "discovery.type=single-node" \
		-e "OPENSEARCH_INITIAL_ADMIN_PASSWORD=$(OS_PSWD)" \
		-v opensearch-data1:/usr/local/data/opensearch \
		opensearchproject/opensearch:latest

# OpenSearch "spelunker" index

os-spelunker-local:
	@make os-spelunker-local-index
	@make os-spelunker-local-fieldlimit

os-spelunker-local-index:
	cat $(WHOSONFIRST_OPENSEARCH)/schema/2.x/mappings.spelunker.json | \
		curl -k \
		-H 'Content-Type: application/json' \
		-X PUT \
		https://admin:$(OS_PSWD)@localhost:9200/spelunker \
		-d @-

os-spelunker-local-fieldlimit:
	cat $(WHOSONFIRST_OPENSEARCH)/schema/2.x/settings.spelunker.json | \
	curl -k \
		-H 'Content-type:application/json' \
		-XPUT \
		https://admin:$(OS_PSWD)@localhost:9200/spelunker/_settings \
		-d @-

# OpenSearch "spelunker" server

os-server-local:
	go run -tags $(GOTAGS_OPENSEARCH) -mod $(GOMOD) ./cmd/wof-spelunker-httpd/main.go \
		-server-uri http://localhost:8080 \
		-spelunker-uri '$(OS_SPELUNKER_URI)'
