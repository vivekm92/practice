package backtracking

import (
	"reflect"
	"testing"
)

func TestNQueens(t *testing.T) {
	tests := []struct {
		A        int
		Expected [][]string
	}{
		{1, [][]string{{"Q"}}},
		{4, [][]string{{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q.."}}},
	}

	for _, test := range tests {
		result := NQueens(test.A)
		if !reflect.DeepEqual(result, test.Expected) {
			t.Errorf("NQueens(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
