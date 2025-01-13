package stringProblems

import "strconv"

/*
	Problem :
*/

// T(n) : O(A * n) ~ O(nlogn), S(n) : O(1)
func CountAndSay(A int) string {

	if A == 1 {
		return "1"
	}

	res := "1"
	for i := 1; i < A; i++ {
		temp := ""
		for j := 0; j < len(res); {
			k, count := j+1, 1
			for k < len(res) && res[k] == res[k-1] {
				count++
				k++
			}
			temp += strconv.Itoa(count) + string(res[k-1])
			j = k
		}
		res = temp
	}

	return res
}
