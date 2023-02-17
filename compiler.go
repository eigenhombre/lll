package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

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

func compile(src string) string {
	// For now, just have LLVM echo the input number.
	n, err := strconv.Atoi(strings.Trim(src, " \n\t"))
	if err != nil {
		panic(err)
	}
	m := ir.NewModule()
	print := m.NewFunc("_print_int", types.Void,
		ir.NewParam("x", types.I32))
	main := m.NewFunc("main", types.I32)
	entry := main.NewBlock("")
	zero := constant.NewInt(types.I32, 0)
	entry.NewCall(print, constant.NewInt(types.I32, int64(n)))
	entry.NewRet(zero)
	return fmt.Sprint(m)
}
