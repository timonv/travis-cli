build: **/*.go test
	go get -v
	go build -v -o bin/travis_cli

install: **/*.go bin/travis_cli build
	cp -f $(CURDIR)/bin/travis_cli /usr/local/bin/travis_cli

test: **/*.go
	go test ./...

clean:
	rm -f bin/travis_cli
	rm -f /usr/local/bin/travis_cli
