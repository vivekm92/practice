package stacksAndQueues

import (
	"testing"
)

func TestRainWaterTrapped(t *testing.T) {
	tests := []struct {
		A        []int
		Expected int
	}{
		{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
		{[]int{4, 2, 0, 3, 2, 5}, 9},
	}

	for _, test := range tests {
		result := RainWaterTrapped(test.A)
		if result != test.Expected {
			t.Errorf("RainWaterTrapped(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}

	for _, test := range tests {
		result := RainWaterTrapped1(test.A)
		if result != test.Expected {
			t.Errorf("RainWaterTrapped1(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}

	// for _, test := range tests {
	// 	result := RainWaterTrapped2(test.A)
	// 	if result != test.Expected {
	// 		t.Errorf("RainWaterTrapped2(%v) = %v ; want %v", test.A, result, test.Expected)
	// 	}
	// }
}
