package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestPushSwap_Integration(t *testing.T) {
	// Test Case 1: Checking simple sequence outputs
	exercise_Cmd := exec.Command("go", "run", "main.go", "2 1 3")
	exercise_Output, exercise_Err := exercise_Cmd.CombinedOutput()
	if exercise_Err != nil {
		t.Fatalf("Binary integration target crash trace: %v", exercise_Err)
	}
	if len(exercise_Output) == 0 {
		t.Errorf("Sorter failed instruction pipeline output generation")
	}
	//test comment
	// Test Case 2: Checking how error formats match specifications
	// Passing a clean duplicate string "1 2 1" to trigger the Error flag
	exercise_ErrCmd := exec.Command("go", "run", "main.go", "1 2 1")
	exercise_ErrOutput, _ := exercise_ErrCmd.CombinedOutput()
	
	// Contains check handles any potential \n variations gracefully
	if !strings.Contains(string(exercise_ErrOutput), "Error") {
		t.Errorf("Duplicate check configuration failure. Expected output to contain 'Error', got: %q", string(exercise_ErrOutput))
	}
}