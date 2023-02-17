.PHONY: all prog test docker readme

all: prog readme test

PROG = lll

${PROG}: *.go
	go build .

test:
	go test -v

docker:
	docker build -t ${PROG} .

readme:
	python updatereadme.py
