package main

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	"gotest.tools/assert"
	"tinygo.org/x/go-llvm"
)

func compileAndRun(s string, t *testing.T) string {
	// Write stdout to a file:
	f, err := os.Create("output.ll")
	assert.NilError(t, err, "could not create output.ll")
	defer f.Close()
	llvmText := compile(s)
	f.Write([]byte(llvmText))
	// fmt.Println("Wrote output.ll")
	// fmt.Print(llvmText)
	// Compile resulting LLVM:
	cmd := exec.Command("clang", "output.ll", "_print.c", "-o", "output")
	stdout, err := cmd.Output()
	assert.NilError(t, err, "clang should not fail")
	assert.DeepEqual(t, string(stdout), "")
	cmd = exec.Command("./output")
	stdout, err = cmd.Output()
	assert.NilError(t, err, "output should not fail")
	return string(stdout)
}

func TestEndToEndIntegration(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"1", "1\n"},
		{"1\n", "1\n"},
		{"  1\n", "1\n"},
		{"2", "2\n"},
		{"99999", "99999\n"},
	}
	// Loop over all the test cases:
	for _, tc := range testCases {
		assert.DeepEqual(t, compileAndRun(tc.input, t), tc.expected)
	}
}

func TestLiveFunctionExecution(t *testing.T) {
	// https://stackoverflow.com/questions/36870139/how-can-jited-llvm-code-call-back-into-a-go-function
	builder := llvm.NewBuilder()
	mod := llvm.NewModule("test")

	// create our function prologue
	main := llvm.FunctionType(llvm.Int32Type(), []llvm.Type{}, false)
	llvm.AddFunction(mod, "main", main)
	block := llvm.AddBasicBlock(mod.NamedFunction("main"), "entry")
	builder.SetInsertPoint(block, block.FirstInstruction())
	builder.CreateRet(llvm.ConstInt(llvm.Int32Type(), 999, false))

	// llvm.LinkInMCJIT()
	// llvm.InitializeNativeTarget()
	// llvm.InitializeNativeAsmPrinter()

	// verify module correctness
	if ok := llvm.VerifyModule(mod, llvm.ReturnStatusAction); ok != nil {
		fmt.Println(ok.Error())
	}
	mod.Dump()

	// create our execution engine
	engine, err := llvm.NewExecutionEngine(mod)
	if err != nil {
		fmt.Println(err.Error())
	}
	// run the function!
	funcResult := engine.RunFunction(mod.NamedFunction("main"), []llvm.GenericValue{})
	result := funcResult.Int(false)
	if result != 999 {
		t.Errorf("Unexpected result, wanted 999, got %d", result)
	}
}
