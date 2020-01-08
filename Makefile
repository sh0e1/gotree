build:
	go build ./cmd/tree

lint:
	@if ! type golint; \
		then go get -u golang.org/x/lint/golint ; \
	fi
	golint -set_exit_status $$(go list ./...)
	go vet ./...

test:
	go test ./...
