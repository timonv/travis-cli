build: **/*.go test
	go get -v
	go build -v -o bin/travis_cli

install soft: **/*.go bin/travis_cli build
	ln -Fs $(CURDIR)/bin/travis_cli /usr/local/bin/travis_cli

install: **/*.go bin/travis_cli build
	cp -f $(CURDIR)/bin/travis_cli /usr/local/bin/travis_cli

test: **/*.go
	go test ./...

clean:
	rm bin/travis_cli
