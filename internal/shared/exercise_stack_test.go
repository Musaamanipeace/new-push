package shared

import "testing"

func exercise_StacksEqual(exercise_S1, exercise_S2 ExerciseStack) bool {
	if len(exercise_S1) != len(exercise_S2) {
		return false
	}
	for exercise_I := range exercise_S1 {
		if exercise_S1[exercise_I] != exercise_S2[exercise_I] {
			return false
		}
	}
	return true
}

func TestExercise_StackMovements(t *testing.T) {
	exercise_A := ExerciseStack{1, 2, 3}
	Exercise_Sa(&exercise_A)
	if !exercise_StacksEqual(exercise_A, ExerciseStack{2, 1, 3}) {
		t.Errorf("Sa operation failed extraction mechanics")
	}

	exercise_B := ExerciseStack{}
	Exercise_Pb(&exercise_A, &exercise_B)
	if len(exercise_B) != 1 || exercise_B[0] != 2 {
		t.Errorf("Pb operation split failure trace")
	}
}