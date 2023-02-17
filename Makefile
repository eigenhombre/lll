.PHONY: all prog test docker

all: prog

PROG = starter

test:
	go test -v

docker:
	docker build -t ${PROG} .
