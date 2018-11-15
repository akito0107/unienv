.PHONY: build
build: tools gen
	go build -o bin/unienv cmd/unienv/main.go

.PHONY: gen
gen:
	go generate ./...

.PHONY: vendor
vendor:
	go mod vendor

clean:
	rm bin/*

tools: bin/generr bin/richgo

bin/generr:
	go build -o bin/generr github.com/akito0107/generr/cmd/generr

bin/richgo:
	go build -o bin/richgo github.com/kyoh86/richgo

.PHONY: test
test: tools
	bin/richgo test -v -cover


.PHONY: e2e
e2e: e2e/unienv
	cd e2e; go test -v .

e2e/unienv:
	go build -o e2e/unienv cmd/unienv/main.go
