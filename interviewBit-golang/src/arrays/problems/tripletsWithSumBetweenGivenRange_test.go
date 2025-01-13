package arrayProblems

import "testing"

func TestTripletsExists(t *testing.T) {
	tests := []struct {
		input  []string
		output int
	}{
		{[]string{"0.297657", "0.420048", "0.053365", "0.437979", "0.375614", "0.264731", "0.060393", "0.197118", "0.203301", "0.128168"}, 1},
		{[]string{"0.6", "0.7", "0.8", "1.2", "0.4"}, 1},
	}

	for _, test := range tests {
		result := TripletsExists(test.input)
		if result != test.output {
			t.Errorf("TripletsExists(%v) = %v ; want %v", test.input, result, test.output)
		}
	}
}
