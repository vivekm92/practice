package arrayProblems

import "testing"

func TestMaxAreaOfTriangle(t *testing.T) {
	tests := []struct {
		input  []string
		output int
	}{
		{[]string{"rrrrr", "rrrrg", "rrrrr", "bbbbb"}, 10},
		{[]string{"rrr", "rrr", "rrr", "rrr"}, 0},
		{[]string{"bbrbbb", "brgrbb", "rbbggb", "rggggr", "rrggrb", "grrbrg", "gbbrrg", "grgrbb", "bbbrgr", "bbrrgg", "rggrbg", "rrgggg", "ggbbgb", "grggbb", "rrrbrr", "rrrbrb", "bbbbbb", "rbbbrg", "ggbbbg", "ggbggr", "bggrgb", "bbrrrb", "rbrrbb", "brbgrg", "rbrrrg", "bbrrgg", "rbgrgg"}, 81},
	}

	for _, test := range tests {
		result := MaxAreaOfTriangle(test.input)
		if result != test.output {
			t.Errorf("MaxAreaOfTriangle(%v) = %v ; want %v", test.input, result, test.output)
		}
	}
}
