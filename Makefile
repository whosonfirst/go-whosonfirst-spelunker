CWD=$(shell pwd)

GOMOD=$(shell test -f "go.work" && echo "readonly" || echo "vendor")

# ./bin/urlencode 'findingaid://https/data.whosonfirst.org/findingaid?template=https://raw.githubusercontent.com/whosonfirst-data/{repo}/main/data/' | pbcopy
READER_URI=findingaid%3A%2F%2Fhttps%2Fdata.whosonfirst.org%2Ffindingaid%3Ftemplate%3Dhttps%3A%2F%2Fraw.githubusercontent.com%2Fwhosonfirst-data%2F%7Brepo%7D%2Fmain%2Fdata%2F

debug:
	go run -mod $(GOMOD) cmd/server/main.go \
		-server-uri http://localhost:8080 \
		-spelunker-uri 'reader://?reader-uri=$(READER_URI)'
