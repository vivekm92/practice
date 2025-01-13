package linkedList

import (
	"fmt"
	"interviewBit/src/utils"
	"testing"
)

type deleteDuplicatesTestCase struct {
	A        *utils.ListNode
	Expected *utils.ListNode
}

func generateDeleteDuplicatesTestCase() []deleteDuplicatesTestCase {
	var deleteDuplicatesTestCases []deleteDuplicatesTestCase
	i1 := utils.GenerateLinkedList([]int{1, 1, 2, 2, 3, 3, 4, 4, 5})
	o1 := utils.GenerateLinkedList([]int{1, 2, 3, 4, 5})
	var t1 = deleteDuplicatesTestCase{
		A:        i1,
		Expected: o1,
	}
	deleteDuplicatesTestCases = append(deleteDuplicatesTestCases, t1)

	i2 := utils.GenerateLinkedList([]int{1, 2, 3, 4, 5})
	o2 := utils.GenerateLinkedList([]int{1, 2, 3, 4, 5})
	var t2 = deleteDuplicatesTestCase{
		A:        i2,
		Expected: o2,
	}
	deleteDuplicatesTestCases = append(deleteDuplicatesTestCases, t2)

	return deleteDuplicatesTestCases
}

func TestDeleteDuplicates(t *testing.T) {
	for idx, test := range generateDeleteDuplicatesTestCase() {
		if output := DeleteDuplicates(test.A); !utils.CompareLinkedLists(output, test.Expected) {
			fmt.Println(test.A.Value, test.Expected.Value, output.Value)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output.Value, test.Expected.Value)
		}
	}
}
