package mathProblems

import (
	"strconv"
)

func NextSimilarNumber(A string) string {

	n := len(A)
	pivot := -1
	var pivotVal int
	for i := n - 2; i >= 0; i-- {
		currIdxVal := A[i] - '0'
		nextIdxVal := A[i+1] - '0'
		if currIdxVal < nextIdxVal {
			pivot = i
			pivotVal = int(currIdxVal)
			break
		}
	}

	if pivot == -1 {
		return "-1"
	}

	lastIdx, lastIdxVal := -1, -1
	for i := n - 1; i > pivot; i-- {
		currIdxVal := int(A[i] - '0')
		if currIdxVal > pivotVal {
			lastIdx = i
			lastIdxVal = currIdxVal
			break
		}
	}

	var temp string
	for i := 0; i < n; i++ {
		if i == pivot {
			temp += strconv.Itoa(lastIdxVal)
		} else if i == lastIdx {
			temp += strconv.Itoa(pivotVal)
		} else {
			temp += string(A[i])
		}
	}

	var res string
	for i := 0; i < pivot+1; i++ {
		res += string(temp[i])
	}
	for i := n - 1; i > pivot; i-- {
		res += string(temp[i])
	}

	return res
}
