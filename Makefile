.PHONY: build test clean build-css dev-css run
POSTCSS=./node_modules/.bin/postcss
CSS_SOURCE=./views/css/tailwind.css
CSS_OUTPUT=public/styles.css

build:
	go build -o myapp

test:
	go test ./...

clean:
	go clean
	rm -f myapp

build-css:
	NODE_ENV=production $(POSTCSS) $(CSS_SOURCE) -o $(CSS_OUTPUT)

dev-css:
	$(POSTCSS) $(CSS_SOURCE) -o $(CSS_OUTPUT)

run:
	air

