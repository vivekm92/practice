package stringProblems

import (
	"strconv"
)

/*
	Problem : https://www.interviewbit.com/problems/string-and-its-frequency/
*/

// T(n) : O(n) , S(n) : O(1)
func StringAndItsFrequency(A string) string {

	arr := make([]int, 26)
	for _, v := range A {
		arr[v-'a']++
	}

	res := ""
	for _, v := range A {
		if arr[v-'a'] > 0 {
			res += string(v) + strconv.Itoa(arr[v-'a'])
			arr[v-'a'] = -1
		}
	}

	return res
}
