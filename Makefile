bin/travis_cli: **/*.go
	@echo "Installing dependencies ..."
	go get -v
	@echo "Running tests ..."
	test
	@echo "Compiling ..."
	go build -v -o bin/travis_cli

install soft: **/*.go bin/travis_cli
	ln -Fs $(CURDIR)/bin/travis_cli /usr/local/bin/travis_cli

install: **/*.go bin/travis_cli
	cp -f $(CURDIR)/bin/travis_cli /usr/local/bin/travis_cli

test: **/*.go
	go test ./...

clean:
	rm bin/travis_cli
