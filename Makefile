.PHONY: all prog docker readme deps

PROG=lll

all: ${PROG}

deps:
	go get .

${PROG}: *.go
	go build .

install: ${PROG}
	go install .

test: ${PROG} .TESTED
	go test -v
	touch .TESTED

docker:
	docker build -t ${PROG} .

readme:
	python updatereadme.py
