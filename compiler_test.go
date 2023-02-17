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
	f.Write([]byte(compile(s)))
	// Compile resulting LLVM:
	cmd := exec.Command("clang", "output.ll", "-o", "output")
	stdout, err := cmd.Output()
	assert.NilError(t, err, "clang should not fail")
	assert.DeepEqual(t, string(stdout), "")
	cmd = exec.Command("./output")
	stdout, err = cmd.Output()
	assert.NilError(t, err, "output should not fail")
	return string(stdout)
}

func TestHelloWorld(t *testing.T) {
	assert.DeepEqual(t, compileAndRun("foo", t), "Hello, world!\n\n")
}
