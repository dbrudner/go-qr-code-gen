.PHONY: build test clean build-css dev-css run

build:
	go build -o myapp

test:
	go test ./...

clean:
	go clean
	rm -f myapp

build-css:
	NODE_ENV=production postcss styles.css -o public/styles.css

dev-css:
	./node_modules/.bin/postcss ./views/css/tailwind.css -o public/styles.css

run:
	air

