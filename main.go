package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"tinygo.org/x/go-llvm"
)

func main() {
	if len(os.Args) != 2 {
		repl()
		fmt.Println()
		os.Exit(0)
	}
	path := os.Args[1]
	src := readFile(path)
	fmt.Print(CompileExternal(src))
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
	bio := bufio.NewReader(os.Stdin)
	// FIXME: don't discard hasMoreInLine
	line, _, err := bio.ReadLine()
	switch err {
	case nil:
		return string(line), nil
	default:
		return "", err
	}
}

func repl() {
	var num int
	var parseTree Expr
	var compiledUnit llvm.Module
	var result int
top:
	fmt.Print("> ")
	s, err := readLine()
	switch err {
	case nil:
	case io.EOF:
		goto done // Spaghetti, anyone?
	default:
		fmt.Println(err)
		goto top
	}
	if s == "" {
		goto top
	}
	parseTree = lexAndParse(s)
	num = parseTree.(Int).Value
	compiledUnit = CompileInternal(num)
	result = ExecInternal(compiledUnit)
	fmt.Println(result)
	goto top
done:
}
