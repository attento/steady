deps:
	go get ./...

tests:
	go test ./...

build:
	go build

clean:
	rm -rf lb_*

compile:
	env GOOS=darwin GOARCH=386 go build
	mv lb lb_$(VERSION)_darwin_386
	env GOOS=linux GOARCH=arm GOARM=7 go build
	mv lb lb_$(VERSION)_linux_arm
	env GOOS=linux GOARCH=386 go build
	mv lb lb_$(VERSION)_linux_386

release: clean compile
