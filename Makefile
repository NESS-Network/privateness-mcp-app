APP=privateness-mcp-app

.PHONY: deps build run docker clean
deps:
	go mod tidy

build: deps
	go build -o bin/$(APP) ./cmd

run: build
	TLS_CERT?=./devcert/cert.pem
	TLS_KEY?=./devcert/key.pem
	BIN=bin/$(APP) TLS_CERT=$$TLS_CERT TLS_KEY=$$TLS_KEY ./bin/$(APP)

docker:
	docker build -t privateness-mcp-app:dev .

clean:
	rm -rf bin
