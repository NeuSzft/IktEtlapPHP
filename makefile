DIR = $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

run:
	docker run --rm -p 8080:8080 -w /go/src/ -v $(DIR):/go/src/ golang:latest go build -o bin/server server/main.go && bin/server

run-detached:
	docker run -d --rm -p 8080:8080 -w /go/src/ -v $(DIR):/go/src/ golang:latest go build -o bin/server server/main.go && bin/server
