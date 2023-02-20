package main

import (
	"os"
	"os/exec"
	"testing"

	"gotest.tools/assert"
)

func compileAndRun(s string, t *testing.T) string {
	// Write stdout to a file:
	f, err := os.Create("output.ll")
	assert.NilError(t, err, "could not create output.ll")
	defer f.Close()
	llvmText := CompileExternal(s)
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
	mod := CompileInternal(999)
	result := ExecInternal(mod)
	if result != 999 {
		t.Errorf("Unexpected result, wanted 999, got %d", result)
	}
}
