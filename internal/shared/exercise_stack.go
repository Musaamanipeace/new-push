package shared

// ExerciseStack acts as our underlying vector tracking numbers across arrays.
type ExerciseStack []int

func Exercise_Pa(exercise_A *ExerciseStack, exercise_B *ExerciseStack) {
	if len(*exercise_B) == 0 {
		return
	}
	exercise_TopElement := (*exercise_B)[0]
	*exercise_B = (*exercise_B)[1:]
	*exercise_A = append(ExerciseStack{exercise_TopElement}, *exercise_A...)
}

func Exercise_Pb(exercise_A *ExerciseStack, exercise_B *ExerciseStack) {
	if len(*exercise_A) == 0 {
		return
	}
	exercise_TopElement := (*exercise_A)[0]
	*exercise_A = (*exercise_A)[1:]
	*exercise_B = append(ExerciseStack{exercise_TopElement}, *exercise_B...)
}

func Exercise_Sa(exercise_A *ExerciseStack) {
	if len(*exercise_A) < 2 {
		return
	}
	(*exercise_A)[0], (*exercise_A)[1] = (*exercise_A)[1], (*exercise_A)[0]
}

func Exercise_Sb(exercise_B *ExerciseStack) {
	if len(*exercise_B) < 2 {
		return
	}
	(*exercise_B)[0], (*exercise_B)[1] = (*exercise_B)[1], (*exercise_B)[0]
}

func Exercise_Ra(exercise_A *ExerciseStack) {
	if len(*exercise_A) < 2 {
		return
	}
	exercise_TopElement := (*exercise_A)[0]
	*exercise_A = append((*exercise_A)[1:], exercise_TopElement)
}

func Exercise_Rb(exercise_B *ExerciseStack) {
	if len(*exercise_B) < 2 {
		return
	}
	exercise_TopElement := (*exercise_B)[0]
	*exercise_B = append((*exercise_B)[1:], exercise_TopElement)
}

func Exercise_Rra(exercise_A *ExerciseStack) {
	if len(*exercise_A) < 2 {
		return
	}
	exercise_Length := len(*exercise_A)
	exercise_BottomElement := (*exercise_A)[exercise_Length-1]
	*exercise_A = append(ExerciseStack{exercise_BottomElement}, (*exercise_A)[:exercise_Length-1]...)
}

func Exercise_Rrb(exercise_B *ExerciseStack) {
	if len(*exercise_B) < 2 {
		return
	}
	exercise_Length := len(*exercise_B)
	exercise_BottomElement := (*exercise_B)[exercise_Length-1]
	*exercise_B = append(ExerciseStack{exercise_BottomElement}, (*exercise_B)[:exercise_Length-1]...)
}