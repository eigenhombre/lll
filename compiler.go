package main

import (
	"fmt"
	"os"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
)

func main() {
	fmt.Print(compile(GetStdin()))
}

// GetStdin reads all of stdin and returns it as a string.
func GetStdin() string {
	s := ""
	for {
		b := make([]byte, 1024)
		n, err := os.Stdin.Read(b)
		if err != nil {
			break
		}
		s += string(b[:n])
	}
	return s
}

func compile(_ string) string {
	// Create a new LLVM IR module.
	m := ir.NewModule()
	hello := constant.NewCharArrayFromString("Hello, world!\n\x00")
	str := m.NewGlobalDef("str", hello)
	// link to external function puts
	puts := m.NewFunc("puts", types.I32,
		ir.NewParam("",
			types.NewPointer(types.I8)))
	main := m.NewFunc("main", types.I32)
	entry := main.NewBlock("")
	zero := constant.NewInt(types.I8, 0)
	// Perform cast per [1]:
	gep := constant.NewGetElementPtr(hello.Typ, str, zero, zero)
	entry.NewCall(puts, gep)
	entry.NewRet(constant.NewInt(types.I32, 0))
	return fmt.Sprint(m)
}

// [1] See also:
// https://github.com/anoopsarkar/compilers-class-hw/blob/master/llvm-practice/helloworld.ll
