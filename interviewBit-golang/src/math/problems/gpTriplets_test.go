package mathProblems

import "testing"

func TestGpTriplets1(t *testing.T) {
	tests := []struct {
		A        []int
		Expected int
	}{
		{[]int{1, 2, 8, 4}, 2},
		{[]int{2, 2, 2, 4}, 6},
		{[]int{2, 2, 2, 2}, 24},
	}
	for _, test := range tests {
		result := GpTriplets1(test.A)
		if result != test.Expected {
			t.Errorf("GpTriplets1(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
