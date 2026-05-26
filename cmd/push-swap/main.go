package main

import (
	"fmt"
	"os"
	"sort"
	"push-swap-project/internal/shared"
)

func exercise_IsSorted(exercise_S shared.ExerciseStack) bool {
	for exercise_I := 0; exercise_I < len(exercise_S)-1; exercise_I++ {
		if exercise_S[exercise_I] > exercise_S[exercise_I+1] {
			return false
		}
	}
	return true
}

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

	if exercise_IsSorted(exercise_StackA) {
		return
	}

	exercise_SortedCopy := append(shared.ExerciseStack{}, exercise_StackA...)
	sort.Ints(exercise_SortedCopy)

	exercise_MapPositions := make(map[int]int)
	for exercise_Idx, exercise_Val := range exercise_SortedCopy {
		exercise_MapPositions[exercise_Val] = exercise_Idx
	}

	if len(exercise_StackA) <= 3 {
		for !exercise_IsSorted(exercise_StackA) {
			if exercise_StackA[0] > exercise_StackA[1] {
				shared.Exercise_Sa(&exercise_StackA)
				fmt.Println("sa")
			} else {
				shared.Exercise_Rra(&exercise_StackA)
				fmt.Println("rra")
			}
		}
		return
	}

	exercise_ChunkSize := 15
	if len(exercise_StackA) > 100 {
		exercise_ChunkSize = 35
	}

	exercise_Counter := 0
	for len(exercise_StackA) > 0 {
		exercise_TargetIndex := exercise_MapPositions[exercise_StackA[0]]
		if exercise_TargetIndex <= exercise_Counter {
			shared.Exercise_Pb(&exercise_StackA, &exercise_StackB)
			fmt.Println("pb")
			exercise_Counter++
		} else if exercise_TargetIndex <= exercise_Counter+exercise_ChunkSize {
			shared.Exercise_Pb(&exercise_StackA, &exercise_StackB)
			fmt.Println("pb")
			shared.Exercise_Rb(&exercise_StackB)
			fmt.Println("rb")
			exercise_Counter++
		} else {
			shared.Exercise_Ra(&exercise_StackA)
			fmt.Println("ra")
		}
	}

	for len(exercise_StackB) > 0 {
		exercise_MaxIdx := 0
		exercise_MaxVal := exercise_StackB[0]
		for exercise_I, exercise_V := range exercise_StackB {
			if exercise_V > exercise_MaxVal {
				exercise_MaxVal = exercise_V
				exercise_MaxIdx = exercise_I
			}
		}

		if exercise_MaxIdx <= len(exercise_StackB)/2 {
			for exercise_I := 0; exercise_I < exercise_MaxIdx; exercise_I++ {
				shared.Exercise_Rb(&exercise_StackB)
				fmt.Println("rb")
			}
		} else {
			for exercise_I := 0; exercise_I < len(exercise_StackB)-exercise_MaxIdx; exercise_I++ {
				shared.Exercise_Rrb(&exercise_StackB)
				fmt.Println("rrb")
			}
		}
		shared.Exercise_Pa(&exercise_StackA, &exercise_StackB)
		fmt.Println("pa")
	}
}