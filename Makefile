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
	rm bin/* e2e/unienv

tools: bin/generr bin/richgo

bin/generr:
	go build -o bin/generr github.com/akito0107/generr/cmd/generr

bin/richgo:
	go build -o bin/richgo github.com/kyoh86/richgo


.PHONY: test
test: test/small test/e2e

.PHONY: test/small
test/small: tools
	bin/richgo test -v -cover


.PHONY: test/e2e
test/e2e: e2e/unienv tools
	cd e2e; ../bin/richgo test -v .

e2e/unienv:
	go build -o e2e/unienv cmd/unienv/main.go
