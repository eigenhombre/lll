package main

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	"gotest.tools/assert"
)

func TestHelloWorld(t *testing.T) {
	// Execute shell command and collect output
	cmd := exec.Command("./starter", "hello")
	stdout, err := cmd.Output()
	assert.NilError(t, err, "starter should not fail")
	fmt.Println(string(stdout))
	// Write stdout to a file:
	f, err := os.Create("output.ll")
	assert.NilError(t, err, "could not create output.ll")
	defer f.Close()
	f.Write(stdout)
	// Compile resulting LLVM:
	cmd = exec.Command("clang", "output.ll", "-o", "output")
	stdout, err = cmd.Output()
	assert.NilError(t, err, "clang should not fail")
	assert.DeepEqual(t, string(stdout), "")
	cmd = exec.Command("./output")
	stdout, err = cmd.Output()
	assert.NilError(t, err, "output should not fail")
	fmt.Println(string(stdout))
}
