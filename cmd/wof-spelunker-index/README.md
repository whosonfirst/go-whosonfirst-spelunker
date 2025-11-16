# wof-spelunker-index

```
$> ./bin/wof-spelunker-index -h
Index one or more Who's On First data sources in a Spelunker-compatible datastore.
Usage: wof-spelunker-index [CMD] [OPTIONS]
Valid commands are:
* opensearch
* sql
```

## Building

The `wof-spelunker-httpd` depends on Go language build tags. The default `cli` Makefile target to compile command line tools build the `wof-spelunker-httpd` tool with support for all the database implementations included in this package. For example:

```
$> cd go-whosonfirst-spelunker
$> make cli
go build -mod vendor -tags="sqlite3,icu,json1,fts5,opensearch" -ldflags="-s -w" -o bin/wof-spelunker-index cmd/wof-spelunker-index/main.go
```

If you only want to build the `wof-spelunker-index` tool with support for SQLite-backed database you can run the `cli-sqlite` Makefile target:

```
$> make cli-sqlite
go build -mod vendor -tags="sqlite3,icu,json1,fts5" -ldflags="-s -w" -o bin/wof-spelunker-index cmd/wof-spelunker-index/main.go
```

_Note that the default SQLite-backed implementation depends on being able to compile the [mattn/go-sqlite3](https://github.com/mattn/go-sqlite3) package._

If you only want to build the `wof-spelunker-index` tool with support for an OpenSearch-backed database you can run the `cli-opensearch` Makefile target:

```
$> make cli-opensearch
go build -mod vendor -tags="opensearch" -ldflags="-s -w" -o bin/wof-spelunker-index cmd/wof-spelunker-index/main.go
```

### Build tags

| Target | Tags | Notes |
| --- | --- | --- |
| MySQL | `mysql` | Support for MySQL should probably still be considered "alpha" at best. |
| Postgres | `postgres` | Support for Postgres should probably still be considered "alpha" at best. |
| SQLite | `sqlite3,icu,json1,fts5` | |
| OpenSearch | `opensearch` | |

## Examples

### database/sql

```
$> ./bin/wof-spelunker-index sql -h
  -database-uri string
    	A URI in the form of 'sql://{DATABASE_SQL_ENGINE}?dsn={DATABASE_SQL_DSN}'. For example: sql://sqlite3?dsn=test.db
  -iterator-uri string
    	A valid whosonfirst/go-whosonfirst-iterate/v3.Iterator URI. Supported iterator URI schemes are: cwd://,directory://,featurecollection://,file://,filelist://,geojsonl://,null://,repo:// (default "repo://")
  -optimize
    	Attempt to optimize the database before closing connection (default true)
  -processes int
    	The number of concurrent processes to index data with (default 28)
  -strict-alt-files
    	Be strict when indexing alt geometries (default true)
  -verbose
    	Enable verbose (debug) logging
```

New `database/sql`-backed Spelunker instances are created by passing a URI to the `NewSpelunker` method in the form of:

```
sql://{DATABASE_ENGINE}?dsn={DATABASE_ENGINE_DSN}
```

Where `{DATABASE_ENGINE}` is a registered (as in "imported") Go language [database/sql](https://pkg.go.dev/database/sql) driver name and `{DATABASE_ENGINE_DSN}` is a driver-specific DSN string for connecting to that database.

See [sql/README.md](../../sql/README.md) for details.

#### SQLite

For example to index all the data in the [whosonfirst-data/whosonfirst-data-admin-ca](#) repository in to a SQLite database called `test.db`:
 ./bin/wof-spelunker-index sql \
	-database-uri 'sql://sqlite3?dsn=test.db' \
	/usr/local/data/whosonfirst/whosonfirst-data-admin-ca/
```

### OpenSearch

```
$> ./bin/wof-spelunker-index opensearch -h
  -create-index
    	Create a new OpenSearch index before indexing records.
  -forgiving
    	 (default true)
  -iterator-uri string
    	A valid whosonfirst/go-whosonfirst-iterate/v3.Iterator URI. Supported iterator URI schemes are: cwd://,directory://,featurecollection://,file://,filelist://,geojsonl://,null://,repo:// (default "repo://")
  -verbose
    	Enable verbose (debug) logging
  -writer-uri string
    	...
```

New OpenSearch-backed Spelunker instances are created by passing a URI to the `NewSpelunker` method in the form of:

```
opensearch://?client_uri={GO_WHOSONFIRST_DATABASE_OPENSEARCH_CLIENT_URI}
```

Where the value of the `client-uri` query parameter is a URL-escaped URI for instantiating a [opensearchapi.Client](https://pkg.go.dev/github.com/opensearch-project/opensearch-go/v4/opensearchapi#Client) instance using the [whosonfirst/go-whosonfirst-database/opensearch/client](https://github.com/whosonfirst/go-whosonfirst-database/tree/main/opensearch/client) package.

For example to index all the data in the [whosonfirst-data/whosonfirst-data-admin-ca](#) repository in to an OpenSearch index named `spelunker`:

```
$> ./bin/wof-spelunker-index opensearch \
	-create-index \
	-writer-uri '...' \
	/usr/local/data/whosonfirst/whosonfirst-data-admin-ca
```

See [opensearch/README.md](../../opensearch/README.md) for details.

