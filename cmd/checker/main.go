package main

import (
	"bufio"
	"fmt"
	"os"
	"push-swap-project/internal/shared"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	exercise_Numbers, exercise_HasError := shared.Exercise_Parser(os.Args[1:])
	if exercise_HasError {
		shared.Exercise_PrintError()
		os.Exit(1)
	}

	var exercise_StackA shared.ExerciseStack = exercise_Numbers
	var exercise_StackB shared.ExerciseStack

	exercise_Scanner := bufio.NewScanner(os.Stdin)
	for exercise_Scanner.Scan() {
		exercise_Cmd := exercise_Scanner.Text()
		if exercise_Cmd == "" {
			continue
		}

		switch exercise_Cmd {
		case "pa": shared.Exercise_Pa(&exercise_StackA, &exercise_StackB)
		case "pb": shared.Exercise_Pb(&exercise_StackA, &exercise_StackB)
		case "sa": shared.Exercise_Sa(&exercise_StackA)
		case "sb": shared.Exercise_Sb(&exercise_StackB)
		case "ra": shared.Exercise_Ra(&exercise_StackA)
		case "rb": shared.Exercise_Rb(&exercise_StackB)
		case "rra": shared.Exercise_Rra(&exercise_StackA)
		case "rrb": shared.Exercise_Rrb(&exercise_StackB)
		default:
			shared.Exercise_PrintError()
			os.Exit(1)
		}
	}

	exercise_IsSorted := true
	for exercise_I := 0; exercise_I < len(exercise_StackA)-1; exercise_I++ {
		if exercise_StackA[exercise_I] > exercise_StackA[exercise_I+1] {
			exercise_IsSorted = false
			break
		}
	}

	if exercise_IsSorted && len(exercise_StackB) == 0 {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}