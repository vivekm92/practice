package linkedList

import (
	"fmt"
	"interviewBit/src/utils"
	"testing"
)

type addTwoNumbersTestCase struct {
	A        *utils.ListNode
	B        *utils.ListNode
	Expected *utils.ListNode
}

func generateAddTwoNumbers() []addTwoNumbersTestCase {
	var addTwoNumbersTestCases []addTwoNumbersTestCase
	ia1 := utils.GenerateLinkedList([]int{2, 4, 3})
	ib1 := utils.GenerateLinkedList([]int{5, 6, 4})
	o1 := utils.GenerateLinkedList([]int{7, 0, 8})
	var t1 = addTwoNumbersTestCase{
		A:        ia1,
		B:        ib1,
		Expected: o1,
	}
	addTwoNumbersTestCases = append(addTwoNumbersTestCases, t1)

	ia2 := utils.GenerateLinkedList([]int{1})
	ib2 := utils.GenerateLinkedList([]int{9, 9, 9})
	o2 := utils.GenerateLinkedList([]int{0, 0, 0, 1})
	var t2 = addTwoNumbersTestCase{
		A:        ia2,
		B:        ib2,
		Expected: o2,
	}
	addTwoNumbersTestCases = append(addTwoNumbersTestCases, t2)

	return addTwoNumbersTestCases
}

func TestAddTwoNumbers(t *testing.T) {
	for idx, test := range generateAddTwoNumbers() {
		if output := AddTwoNumbers(test.A, test.B); !utils.CompareLinkedLists(output, test.Expected) {
			fmt.Println(test.A.Value, test.B.Value, test.Expected.Value, output.Value)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output.Value, test.Expected.Value)
		}
	}
}
