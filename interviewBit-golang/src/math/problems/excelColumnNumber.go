package mathProblems

import "testing"

// T(n) : O(n), S(n) : O(1)
func TitleToNumber(A string) int {

	n, res, pow := len(A), 0, 1
	for i := n - 1; i >= 0; i-- {
		res += pow * (int(A[i]) - int('A') + 1)
		pow *= 26
	}

	return res
}

func TestTitleToNumber(t *testing.T) {

}
