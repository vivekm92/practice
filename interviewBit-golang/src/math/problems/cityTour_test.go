package mathProblems

import "testing"

func TestCityTour(t *testing.T) {
	tests := []struct {
		A        int
		B        []int
		Expected int
	}{
		{3, []int{1}, 1},
		{8, []int{3, 6}, 180},
		{7, []int{4, 3, 2, 5, 6, 7}, 1},
	}

	for _, test := range tests {
		result := CityTour(test.A, test.B)
		if result != test.Expected {
			t.Errorf("CityTour(%v, %v) = %v ; want %v", test.A, test.B, result, test.Expected)
		}
	}
}
