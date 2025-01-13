package bsearch

import "testing"

func TestKthsmallest(t *testing.T) {
	tests := []struct {
		A        []int
		B        int
		Expected int
	}{
		{[]int{2, 1, 4, 3, 2}, 3, 2},
		{[]int{1, 2}, 2, 2},
		{[]int{18, 23, 2, 2, 4, 5, 9}, 7, 23},
	}

	for _, test := range tests {
		result := Kthsmallest(test.A, test.B)
		if result != test.Expected {
			t.Errorf("Kthsmallest(%v, %v) = %v ; want %v", test.A, test.B, result, test.Expected)
		}
	}
}
