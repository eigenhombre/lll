.PHONY: all prog

all: prog

PROG = starter

prog: ${PROG}.go
	go build -o ${PROG} ${PROG}.go

test: prog
	go test -v