package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		repl()
	}
	path := os.Args[1]
	src := readFile(path)
	fmt.Print(compile(src))
}

func readFile(path string) string {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func readLine() string {
	var s string
	_, err := fmt.Scanln(&s)
	if err != nil {
		panic(err)
	}
	return s
}

func repl() {
top:
	fmt.Print("> ")
	for {
		s := readLine()
		if s == "" {
			goto top
		}
		fmt.Print(compile(s))
	}
}
