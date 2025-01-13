package arrayProblems

import (
	"sort"
	"strings"
)

// T(n) : O(n), S(n) : O(1), where n is length of string
func IsDigitLogs(s string) bool {

	idx := strings.Index(s, "-")
	for i := idx; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			return true
		}
	}

	return false
}

// T(n) : O(nlogn), S(n) : O(n)
func ReorderLogs(A []string) []string {

	var digitLogs []string
	var sortedLogs []string
	for _, v := range A {
		if IsDigitLogs(v) {
			digitLogs = append(digitLogs, v)
		} else {
			sortedLogs = append(sortedLogs, v)
		}
	}

	sort.Slice(sortedLogs, func(i, j int) bool {
		i1, j1 := strings.Index(sortedLogs[i], "-"), strings.Index(sortedLogs[j], "-")
		s1, s2 := sortedLogs[i][i1+1:], sortedLogs[j][j1+1:]
		s3, s4 := sortedLogs[i][:i1], sortedLogs[j][:j1]
		if strings.Compare(s1, s2) == 0 {
			return strings.Compare(s3, s4) == -1
		}

		return strings.Compare(s1, s2) == -1
	})

	sortedLogs = append(sortedLogs, digitLogs...)
	return sortedLogs
}
