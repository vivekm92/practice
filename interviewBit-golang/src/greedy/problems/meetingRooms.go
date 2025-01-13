package greedy

import (
	"sort"
)

/*
	Problem : https://www.interviewbit.com/problems/meeting-rooms/
*/

// T(n) : O(nlogn) ; S(n) : O(n)
func MeetingRooms(A [][]int) int {

	arr := make([][]int, 0)
	for _, v := range A {
		arr = append(arr, []int{v[0], 1})
		arr = append(arr, []int{v[1], -1})
	}

	sort.Slice(arr, func(i, j int) bool {
		if arr[i][0] == arr[j][0] {
			return arr[i][1] < arr[j][1]
		}
		return arr[i][0] < arr[j][0]
	})

	cnt, maxCnt := 0, 0
	for _, v := range arr {
		if v[1] == 1 {
			cnt++
		} else {
			cnt--
		}

		if cnt > maxCnt {
			maxCnt = cnt
		}
	}

	return maxCnt
}
