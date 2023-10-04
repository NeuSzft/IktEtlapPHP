EXEPATH = ./bin/server

build:
	go build -o ${EXEPATH} server/main.go

run: build
	${EXEPATH}
