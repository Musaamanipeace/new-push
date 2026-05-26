package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestPushSwap_Integration(t *testing.T) {
	exercise_Cmd := exec.Command("go", "run", "main.go", "2 1 3")
	exercise_Output, exercise_Err := exercise_Cmd.CombinedOutput()
	if exercise_Err != nil {
		t.Fatalf("Binary integration target crash trace: %v", exercise_Err)
	}
	if len(exercise_Output) == 0 {
		t.Errorf("Sorter failed instruction pipeline output generation")
	}

	exercise_ErrCmd := exec.Command("go", "run", "main.go", "1 2 1")
	exercise_ErrOutput, _ := exercise_ErrCmd.CombinedOutput()
	if strings.TrimSpace(string(exercise_ErrOutput)) != "Error" {
		t.Errorf("Duplicate check configuration failure")
	}
}