package greedy

import (
	"testing"
)

func TestSeats(t *testing.T) {

	tests := []struct {
		A        string
		Expected int
	}{
		{".x..x..x.", 4},
	}

	for _, test := range tests {
		result := Seats(test.A)
		if result != test.Expected {
			t.Errorf("Seats(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
