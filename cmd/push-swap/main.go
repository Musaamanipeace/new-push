package main

import (
	"fmt"
	"os"
	"sort"
	"push-swap-project/internal/shared"
)

func exercise_IsSorted(exercise_S shared.ExerciseStack) bool {
	for i := 0; i < len(exercise_S)-1; i++ {
		if exercise_S[i] > exercise_S[i+1] {
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

	// Structural Index Mapping (Transforms raw numbers into values 0 to N-1)
	exercise_SortedCopy := append(shared.ExerciseStack{}, exercise_StackA...)
	sort.Ints(exercise_SortedCopy)

	exercise_MapPositions := make(map[int]int)
	for idx, val := range exercise_SortedCopy {
		exercise_MapPositions[val] = idx
	}

	for idx, val := range exercise_StackA {
		exercise_StackA[idx] = exercise_MapPositions[val]
	}

	// ---------------------------------------------------------
	// PURE NON-COMPARATIVE SMALL PATH (For <= 6 elements)
	// ---------------------------------------------------------
	if len(exercise_StackA) <= 6 {
		totalElements := len(exercise_StackA)
		
		// Sort by pulling elements out sequentially based on their absolute index identity
		for target := 0; target < totalElements; target++ {
			// Find where our target item is sitting structurally in stack A
			targetIdx := -1
			for idx, val := range exercise_StackA {
				if val == target {
					targetIdx = idx
					break
				}
			}

			// Bring it to the top using the shortest rotational path
			if targetIdx <= len(exercise_StackA)/2 {
				for i := 0; i < targetIdx; i++ {
					shared.Exercise_Ra(&exercise_StackA)
					fmt.Println("ra")
				}
			} else {
				for i := 0; i < len(exercise_StackA)-targetIdx; i++ {
					shared.Exercise_Rra(&exercise_StackA)
					fmt.Println("rra")
				}
			}

			// Push to B once it reaches the top
			shared.Exercise_Pb(&exercise_StackA, &exercise_StackB)
			fmt.Println("pb")
		}

		// Stack B is now in perfect descending order; push everything back to A
		for len(exercise_StackB) > 0 {
			shared.Exercise_Pa(&exercise_StackA, &exercise_StackB)
			fmt.Println("pa")
		}
		return
	}

	// ---------------------------------------------------------
	// LARGE DATASET PATH: Multi-Pass Radix Sort (For 100 elements)
	// ---------------------------------------------------------
	exercise_TotalSize := len(exercise_StackA)
	exercise_MaxBits := 0
	for ((exercise_TotalSize - 1) >> exercise_MaxBits) > 0 {
		exercise_MaxBits++
	}

	for bit := 0; bit < exercise_MaxBits; bit++ {
		if exercise_IsSorted(exercise_StackA) && len(exercise_StackB) == 0 {
			break
		}

		for i := 0; i < exercise_TotalSize; i++ {
			exercise_TopElement := exercise_StackA[0]
			if ((exercise_TopElement >> bit) & 1) == 1 {
				shared.Exercise_Ra(&exercise_StackA)
				fmt.Println("ra")
			} else {
				shared.Exercise_Pb(&exercise_StackA, &exercise_StackB)
				fmt.Println("pb")
			}
		}

		for len(exercise_StackB) > 0 {
			shared.Exercise_Pa(&exercise_StackA, &exercise_StackB)
			fmt.Println("pa")
		}
	}
}