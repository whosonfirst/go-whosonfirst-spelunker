CWD=$(shell pwd)

GOMOD=$(shell test -f "go.work" && echo "readonly" || echo "vendor")

# ./bin/urlencode 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:122.0) Gecko/20100101 Firefox/122.0' | pbcopy
# ./bin/urlencode 'findingaid://https/data.whosonfirst.org/findingaid?&user-agent=Mozilla%2F5.0+%28Macintosh%3B+Intel+Mac+OS+X+10.15%3B+rv%3A122.0%29+Gecko%2F20100101+Firefox%2F122.0' | pbcopy

READER_URI=findingaid%3A%2F%2Fhttps%2Fdata.whosonfirst.org%2Ffindingaid%3F%26user-agent%3DMozilla%252F5.0%2B%2528Macintosh%253B%2BIntel%2BMac%2BOS%2BX%2B10.15%253B%2Brv%253A122.0%2529%2BGecko%252F20100101%2BFirefox%252F122.0

debug:
	go run -mod $(GOMOD) cmd/server/main.go \
		-server-uri http://localhost:8080 \
		-spelunker-uri 'reader://?reader-uri=$(READER_URI)'
