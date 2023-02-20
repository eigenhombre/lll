package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"tinygo.org/x/go-llvm"
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

func CompileExternal(src string) string {
	// For now, just have LLVM echo the input number.
	// FIXME: Change this to use go-llvm and link to
	// _print_int or some Go output function.
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

// c.f.: https://stackoverflow.com/questions/36870139/
func CompileInternal(arg int) llvm.Module {
	builder := llvm.NewBuilder()
	mod := llvm.NewModule("test")
	main := llvm.FunctionType(llvm.Int32Type(), []llvm.Type{}, false)
	llvm.AddFunction(mod, "main", main)
	block := llvm.AddBasicBlock(mod.NamedFunction("main"), "entry")
	builder.SetInsertPoint(block, block.FirstInstruction())
	y := uint64(arg)
	builder.CreateRet(llvm.ConstInt(llvm.Int32Type(), y, false))
	return mod
}

func ExecInternal(mod llvm.Module) int {
	// verify module correctness
	if ok := llvm.VerifyModule(mod, llvm.ReturnStatusAction); ok != nil {
		fmt.Println(ok.Error())
	}
	// mod.Dump()
	// create our execution engine
	engine, err := llvm.NewExecutionEngine(mod)
	if err != nil {
		fmt.Println(err.Error())
	}
	// run the function!
	funcResult := engine.RunFunction(mod.NamedFunction("main"), []llvm.GenericValue{})
	result := funcResult.Int(false)
	return int(result)
}
