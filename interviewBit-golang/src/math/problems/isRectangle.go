package mathProblems

import "testing"

// T(n) : O(1), S(n) : O(1)
func IsRectangle(A int, B int, C int, D int) int {

	lookup := make(map[int]int)

	lookup[A] += 1
	lookup[B] += 1
	lookup[C] += 1
	lookup[D] += 1

	if len(lookup) == 1 {
		return 1
	} else if len(lookup) == 2 {
		for _, v := range lookup {
			if v != 2 {
				return 0
			}
		}
		return 1
	}

	return 0
}

func TestIsRectangle(t *testing.T) {

}
