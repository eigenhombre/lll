package main

import (
	"fmt"
	"os"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
)

func hello() {
	// Create a new LLVM IR module.
	m := ir.NewModule()
	arr := constant.NewCharArrayFromString("Hello, world!\n\x00")
	gd := m.NewGlobalDef("str", arr)
	// link to external function puts
	puts := m.NewFunc("puts", types.I32, ir.NewParam("", types.NewPointer(types.I8)))
	main := m.NewFunc("main", types.I32)
	entry := main.NewBlock("")

	// Perform cast per [1]:
	gep := constant.NewGetElementPtr(types.NewArray(15, types.I8),
		gd,
		constant.NewInt(types.I8, 0),
		constant.NewInt(types.I8, 0))
	entry.NewCall(puts, gep)
	entry.NewRet(constant.NewInt(types.I32, 0))
	fmt.Println(m)
}

// Hello, World in LLVM IR.
func main() {
	// get first argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: starter <progname>")
		os.Exit(1)
	}
	// switch on argument
	switch os.Args[1] {
	case "hello":
		hello()
	default:
		fmt.Println("Unknown program name")
		os.Exit(1)
	}
}

// [1] See also:
// https://github.com/anoopsarkar/compilers-class-hw/blob/master/llvm-practice/helloworld.ll
