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

tools: bin/generr

bin/generr:
	go build -o bin/generr github.com/akito0107/generr/cmd/generr
