generate: gen/node_modules
	rm -f html_*.go
	node gen/gen.js
	gofmt -w .

gen/node_modules:
	cd gen; npm install

catalog: gen/spec/catalog_src.html
	node gen/spec/gen_catalog.js
