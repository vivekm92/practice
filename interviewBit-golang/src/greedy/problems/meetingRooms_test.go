package greedy

import "testing"

func TestMeetingRooms(t *testing.T) {

	tests := []struct {
		A        [][]int
		Expected int
	}{
		{[][]int{{1, 2}, {3, 4}, {5, 6}}, 1},
		{[][]int{{1, 4}, {3, 4}, {5, 6}}, 2},
	}

	for _, test := range tests {
		result := MeetingRooms(test.A)
		if result != test.Expected {
			t.Errorf("MeetingRooms(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
