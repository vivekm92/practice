package stacksAndQueues

import (
	"reflect"
	"strconv"
	"testing"
)

// Driver Code for Testing
func RunMinStack(A [][]string) []int {

	res := make([]int, 0)
	obj := MinStackConstructor()
	for _, v := range A {
		if v[0] == "Push" {
			val, _ := strconv.Atoi(v[1])
			obj.Push(val)
		} else if v[0] == "Pop" {
			obj.Pop()
		} else if v[0] == "Top" {
			val := obj.Top()
			res = append(res, val)
		} else if v[0] == "GetMin" {
			val := obj.GetMin()
			res = append(res, val)
		}
	}

	return res
}

func TestMinStack(t *testing.T) {
	tests := []struct {
		A        [][]string
		Expected []int
	}{
		{[][]string{
			{"Push", "-2"},
			{"Push", "0"},
			{"Push", "-3"},
			{"GetMin", ""},
			{"Pop", ""},
			{"Top", ""},
			{"GetMin", ""}}, []int{-3, 0, -2}},
	}

	for _, test := range tests {
		result := RunMinStack(test.A)
		if !reflect.DeepEqual(result, test.Expected) {
			t.Errorf("RunMinStack(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
