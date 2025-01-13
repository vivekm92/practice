package mathProblems

import "testing"

func TestSocksPairs(t *testing.T) {
	tests := []struct {
		A        []int
		Expected int
	}{
		{[]int{1, 2, 3}, 0},
		{[]int{2, 2, 2, 2}, 2},
	}

	for _, test := range tests {
		result := SocksPairs(test.A)
		if result != test.Expected {
			t.Errorf("SocksPairs(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
