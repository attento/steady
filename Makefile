.PHONY: deps tests build clean compile

tests:
	go test ./...

build: dist/steady

dist/steady:
	mkdir -p dist/
	cd dist/; go build -o steady ..

clean:
	rm -rf dist/

compile:
	mkdir -p dist/
	cd dist/; \
		env GOOS=darwin GOARCH=386 go build -o steady_$(VERSION)_darwin_386 .. \
	    env GOOS=linux GOARCH=arm go build -o steady_$(VERSION)_linux_arm .. \
		env GOOS=linux GOARCH=arm64 go build -o steady_$(VERSION)_linux_arm64 .. \
		env GOOS=linux GOARCH=386 go build -o steady_$(VERSION)_linux_386 ..


release: clean compile
