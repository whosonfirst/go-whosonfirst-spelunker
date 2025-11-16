# wof-spelunker-httpd

Start the Spelunker web application.

```
$> ./bin/wof-spelunker-httpd -h
Start the Spelunker web application.
Usage:
	./bin/wof-spelunker-httpd [options]
Valid options are:
  -authenticator-uri string
    	A valid aaronland/go-http/v3/auth.Authenticator URI. This is future-facing work and can be ignored for now. (default "null://")
  -map-provider string
    	Valid options are: leaflet, protomaps (default "leaflet")
  -map-tile-uri string
    	A valid Leaflet tile layer URI. See documentation for special-case (interpolated tile) URIs. (default "https://tile.openstreetmap.org/{z}/{x}/{y}.png")
  -protomaps-max-data-zoom int
    	The maximum zoom (tile) level for data in a PMTiles database
  -protomaps-theme string
    	A valid Protomaps theme label. (default "white")
  -root-url string
    	The root URL for all public-facing URLs and links. If empty then the value of the -server-uri flag will be used.
  -server-uri string
    	A valid `aaronland/go-http/v3/server.Server URI. (default "http://localhost:8080")
  -spelunker-uri string
    	A URI in the form of '{SPELUNKER_SCHEME}://{IMPLEMENTATION_DETAILS}' referencing the underlying Spelunker database. For example: sql://sqlite3?dsn=spelunker.db (default "null://")
```

## Building

The `wof-spelunker-httpd` depends on Go language build tags. The default `cli` Makefile target to compile command line tools build the `wof-spelunker-httpd` tool with support for all the database implementations included in this package. For example:

```
$> cd go-whosonfirst-spelunker
$> make cli
go build -mod vendor -tags="sqlite3,icu,json1,fts5,opensearch" -ldflags="-s -w" -o bin/wof-spelunker-httpd cmd/wof-spelunker-httpd/main.go
```

If you only want to build the `wof-spelunker-httpd` tool with support for SQLite-backed database you can run the `cli-sqlite` Makefile target:

```
$> make cli-sqlite
go build -mod vendor -tags="sqlite3,icu,json1,fts5" -ldflags="-s -w" -o bin/wof-spelunker-httpd cmd/wof-spelunker-httpd/main.go
```

_Note that the default SQLite-backed implementation depends on being able to compile the [mattn/go-sqlite3](https://github.com/mattn/go-sqlite3) package._

If you only want to build the `wof-spelunker-httpd` tool with support for an OpenSearch-backed database you can run the `cli-opensearch` Makefile target:

```
$> make cli-opensearch
go build -mod vendor -tags="opensearch" -ldflags="-s -w" -o bin/wof-spelunker-httpd cmd/wof-spelunker-httpd/main.go
```

### Build tags

| Target | Tags | Notes |
| --- | --- | --- |
| MySQL | `mysql` | Support for MySQL should probably still be considered "alpha" at best. |
| Postgres | `postgres` | Support for Postgres should probably still be considered "alpha" at best. |
| SQLite | `sqlite3,icu,json1,fts5` | |
| OpenSearch | `opensearch` | |

## Maps

Map configuration for the `wof-spelunker-httpd` application is controlled by these flags:

```
  -map-provider string
    	Valid options are: leaflet, protomaps (default "leaflet")
  -map-tile-uri string
    	A valid Leaflet tile layer URI. See documentation for special-case (interpolated tile) URIs. (default "https://tile.openstreetmap.org/{z}/{x}/{y}.png")
  -protomaps-max-data-zoom int
    	The maximum zoom (tile) level for data in a PMTiles database (default 15)
  -protomaps-theme string
    	A valid Protomaps theme label. (default "white")
```

_Under the hood the `wof-spelunker-httpd` application is using the [aaronland/go-http-maps](https://github.com/aaronland/go-http-maps) package to manage map configuration._

### Raster tiles

The default map for the `wof-spelunker-httpd` application uses the [LeafletJS](https://leafletjs.com/) package to display raster tiles provided by the [OpenStreetMap](https://) project.

You can use alternate raster tiles by specifying their "ZXY" tile URL in the `-map-file-uri` flag.

### Protomaps tiles

You can use also use map data encoded in [PMTiles](https://docs.protomaps.com/pmtiles/) database for rendering base maps by setting the `-map-provider` flag to "protomaps".

The value of the `-map-tile-uri` should be one of the following:

* `api://{PROTOMAPS_API_KEY}` to load PMTiles data from the [Protomaps API](https://protomaps.com/api).
* `file:///path/to/local/pmtiles.db` to load PMTiles data from a local PMTile database file.

## Examples

### database/sql

New `database/sql`-backed Spelunker instances are created by passing a URI to the `NewSpelunker` method in the form of:

```
sql://{DATABASE_ENGINE}?dsn={DATABASE_ENGINE_DSN}
```

Where `{DATABASE_ENGINE}` is a registered (as in "imported") Go language [database/sql](https://pkg.go.dev/database/sql) driver name and `{DATABASE_ENGINE_DSN}` is a driver-specific DSN string for connecting to that database.

See [sql/README.md](../../sql/README.md) for details.

#### SQLite

For example to start the `wof-spelunker-httpd` application using data stored in a local SQLite database:

```
./bin/wof-spelunker-httpd \
	-spelunker-uri 'sql://sqlite3?dsn=/usr/local/data/sfom.db'
```

### OpenSearch

New OpenSearch-backed Spelunker instances are created by passing a URI to the `NewSpelunker` method in the form of:

```
opensearch://?client_uri={GO_WHOSONFIRST_DATABASE_OPENSEARCH_CLIENT_URI}
```

Where the value of the `client-uri` query parameter is a URL-escaped URI for instantiating a [opensearchapi.Client](https://pkg.go.dev/github.com/opensearch-project/opensearch-go/v4/opensearchapi#Client) instance using the [whosonfirst/go-whosonfirst-database/opensearch/client](https://github.com/whosonfirst/go-whosonfirst-database/tree/main/opensearch/client) package.

For example:

```
./bin/wof-spelunker-httpd \
	-spelunker-uri 'opensearch://?client_uri=https%3A%2F%2Flocalhost%3A9200%3Findex%3Dspelunker%26require-tls%3Dtrue'
```

_The need to URL-escape the `client-uri` parameter is not great but that's how things work today._

See [opensearch/README.md](../../opensearch/README.md) for details.

