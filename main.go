package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		repl()
		fmt.Println()
		os.Exit(0)
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

func readLine() (string, error) {
	var s string
	_, err := fmt.Scanln(&s)

	if err != nil {
		return "", err
	}
	return s, nil
}

func repl() {
top:
	fmt.Print("> ")
	s, err := readLine()
	switch err {
	case nil:
	case io.EOF:
		goto done // Spaghetti, anyone?
	default:
		panic(err)
	}
	if s == "" {
		goto top
	}
	fmt.Print(compile(s))
	goto top
done:
}
