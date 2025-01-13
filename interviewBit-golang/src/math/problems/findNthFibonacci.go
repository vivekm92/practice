// refer : https://medium.com/codebrace/starred-problem-2-nth-fibonacci-number-in-log-n-time-821ea9a18296
package mathProblems

import (
	"testing"
)

func bar(arr1 []int, arr2 []int) []int {

	res, modv := []int{0, 0, 0, 0}, 1000000007
	for i := 0; i < 4; i++ {
		arr1[i] = arr1[i] % modv
		arr2[i] = arr2[i] % modv
	}

	res[0] = ((arr1[0]*arr2[0])%modv + (arr1[1]*arr2[2])%modv) % modv
	res[1] = ((arr1[0]*arr2[1])%modv + (arr1[1]*arr2[3])%modv) % modv
	res[2] = ((arr1[2]*arr2[0])%modv + (arr1[3]*arr2[2])%modv) % modv
	res[3] = ((arr1[2]*arr2[1])%modv + (arr1[3]*arr2[3])%modv) % modv

	return res
}

func foo(arr []int, n int) []int {
	if n == 1 {
		return arr
	}

	temp := foo(arr, n/2)
	if n%2 == 0 {
		return bar(temp, temp)
	}

	return bar(bar(temp, temp), arr)
}

// T(n) : O(log(n)), S(n) : O(1)
func FindNthFibonacci(A int) int {

	if A == 1 || A == 2 {
		return 1
	}

	arr, modv := []int{1, 1, 1, 0}, 1000000007
	arr = foo(arr, A-1)
	return arr[0] % modv
}

func TestFindNthFibonacci(t *testing.T) {

}
