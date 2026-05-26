package main

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"
)

func TestChecker_Integration(t *testing.T) {
	exercise_Cmd := exec.Command("go", "run", "main.go", "3 2 1 0")
	var exercise_Buf bytes.Buffer
	exercise_Buf.WriteString("rra\npb\n")
	exercise_Cmd.Stdin = &exercise_Buf

	exercise_Output, _ := exercise_Cmd.CombinedOutput()
	if !strings.Contains(string(exercise_Output), "KO") && !strings.Contains(string(exercise_Output), "OK") {
		t.Errorf("Verification loop state failure trace")
	}
}