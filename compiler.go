package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
)

type Int struct {
	Value int
}

type Atom struct {
	Name string
}

type Expr interface {
	// Tighten this up?
}

func lexAndParse(src string) Expr {
	n, err := strconv.Atoi(strings.Trim(src, " \n\t"))
	if err != nil {
		panic(err)
	}
	return Int{n}
}

func compile(src string) string {
	// For now, just have LLVM echo the input number.
	ast := lexAndParse(src).(Int).Value
	m := ir.NewModule()
	print := m.NewFunc("_print_int", types.Void,
		ir.NewParam("x", types.I32))
	main := m.NewFunc("main", types.I32)
	entry := main.NewBlock("")
	zero := constant.NewInt(types.I32, 0)
	entry.NewCall(print, constant.NewInt(types.I32, int64(ast)))
	entry.NewRet(zero)
	return fmt.Sprint(m)
}
