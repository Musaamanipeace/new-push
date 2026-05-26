package shared

import "testing"

func TestExercise_Parser(t *testing.T) {
	res, err := Exercise_Parser([]string{"5 67 3"})
	if err || len(res) != 3 || res[1] != 67 {
		t.Errorf("Failed multi-string space isolation block split test")
	}

	_, err = Exercise_Parser([]string{"0", "one", "2"})
	if !err {
		t.Errorf("Failed validation flag for textual input strings")
	}
}