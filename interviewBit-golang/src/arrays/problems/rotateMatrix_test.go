package arrayProblems

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {

	tests := []struct {
		A        [][]int
		Expected [][]int
	}{
		{
			[][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15, 16},
			},
			[][]int{
				{13, 9, 5, 1},
				{14, 10, 6, 2},
				{15, 11, 7, 3},
				{16, 12, 8, 4},
			},
		},
		{
			[][]int{
				{1, 2},
				{3, 4},
			},
			[][]int{
				{3, 1},
				{4, 2},
			},
		},
		{
			[][]int{
				{1},
			},
			[][]int{
				{1},
			},
		},
	}

	for _, test := range tests {
		result := Rotate(test.A)
		if !reflect.DeepEqual(result, test.Expected) {
			t.Errorf("Rotate(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}

func TestRotateMatric(t *testing.T) {

	tests := []struct {
		A        [][]int
		Expected [][]int
	}{
		{
			[][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15, 16},
			},
			[][]int{
				{13, 9, 5, 1},
				{14, 10, 6, 2},
				{15, 11, 7, 3},
				{16, 12, 8, 4},
			},
		},
		{
			[][]int{
				{1, 2},
				{3, 4},
			},
			[][]int{
				{3, 1},
				{4, 2},
			},
		},
		{
			[][]int{
				{1},
			},
			[][]int{
				{1},
			},
		},
	}

	for _, test := range tests {
		result := RotateMatrix(test.A)
		if !reflect.DeepEqual(result, test.Expected) {
			t.Errorf("rotateMatrix(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
