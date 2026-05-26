package shared

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Exercise_Parser reads argument vectors to cleanly isolate and store clean numerical arrays.
func Exercise_Parser(exercise_Args []string) ([]int, bool) {
	if len(exercise_Args) == 0 {
		return nil, false
	}

	var exercise_RawNumbers []string
	if len(exercise_Args) == 1 {
		exercise_RawNumbers = strings.Fields(exercise_Args[0])
	} else {
		exercise_RawNumbers = exercise_Args
	}

	if len(exercise_RawNumbers) == 0 {
		return nil, false
	}

	var exercise_StoredValues []int
	exercise_SeenMap := make(map[int]bool)

	for _, exercise_Token := range exercise_RawNumbers {
		exercise_Num, exercise_Err := strconv.Atoi(exercise_Token)
		if exercise_Err != nil {
			return nil, true 
		}
		if exercise_SeenMap[exercise_Num] {
			return nil, true 
		}
		exercise_SeenMap[exercise_Num] = true
		exercise_StoredValues = append(exercise_StoredValues, exercise_Num)
	}
	return exercise_StoredValues, false
}

func Exercise_PrintError() {
	fmt.Fprintln(os.Stderr, "Error")
}