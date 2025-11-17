# opensearch

The `opensearch` package implements the `Spelunker` interface for Who's On First data indexed in an [OpenSearch](https://opensearch.org/) database.

## Examples

### Running locally

These examples assume a "local" setup meaning there is local instance of OpenSearch running on port 9200. The "easiest" way to do this is with the [Docker](https://www.docker.com/) application running a containerized instance of OpenSearch and the `os-local` Makefile target provided by this package.

For example, in one terminal window:

```
$> cd go-whosonfirst-spelunker
$> make os-local
docker run \
		-it \
		-p 9200:9200 \
		-p 9600:9600 \
		-e "discovery.type=single-node" \
		-e "OPENSEARCH_INITIAL_ADMIN_PASSWORD=dkjfhsjdkfkjdjhksfhskd98475kjHkzjxckj" \
		-v opensearch-data1:/usr/local/data/opensearch \
		opensearchproject/opensearch:latest

...wait for Docker/OpenSearch to start
```

In another terminal run the `os-index-local` Makefile target:

```
$> make os-index-local REPOS=/usr/local/data/whosonfirst/whosonfirst-data-admin-ca
go run -tags opensearch -mod readonly ./cmd/wof-spelunker-index/main.go opensearch \
		-database-uri 'opensearch2://localhost:9200/spelunker?username=admin&password=dkjfhsjdkfkjdjhksfhskd98475kjHkzjxckj&insecure=true&require-tls=true' \
		/usr/local/data/whosonfirst/whosonfirst-data-admin-ca

2025/11/15 17:34:59 INFO Iterator stats elapsed=17.295467917s seen=33845 allocated="229 MB" "total allocated"="15 GB" sys="643 MB" numgc=182
2025/11/15 17:35:08 INFO Index complete indexed=28097
```

Once complete run the `os-server-local` Makefile target:

```
$> make os-server-local
go run -tags opensearch -mod vendor ./cmd/wof-spelunker-httpd/main.go \
		-server-uri http://localhost:8080 \
		-spelunker-uri 'opensearch://?client-uri=https%3A%2F%2Flocalhost%3A9200%2Fspelunker%3Fusername%3Dadmin%26password%3Ddkjfhsjdkfkjdjhksfhskd98475kjHkzjxckj%26insecure%3Dtrue%26require-tls%3Dtrue&cache-uri=ristretto%3A%2F%2F&reader-uri=https%3A%2F%2Fdata.whosonfirst.org'

2025/11/15 11:42:44 INFO Listening for requests address=http://localhost:8080
```