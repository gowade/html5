generate: gen/node_modules
	rm -f html_*.go
	node gen/gen.js
	gofmt -w .

gen/node_modules:
	cd gen; npm install
