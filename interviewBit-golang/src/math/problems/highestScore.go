package mathProblems

import (
	"strconv"
)

type B struct {
	score, count int
}

// T(n) : O(n), S(n) : O(n)
func HighestScore(A [][]string) int {
	lookup := make(map[string]*B)
	for _, v := range A {
		val, _ := strconv.Atoi(v[1])
		if _, ok := lookup[v[0]]; ok {
			lookup[v[0]].score = lookup[v[0]].score + val
			lookup[v[0]].count = lookup[v[0]].count + 1
		} else {
			lookup[v[0]] = new(B)
			lookup[v[0]].score = val
			lookup[v[0]].count = 1
		}
	}

	var maxAverage int
	for _, v := range lookup {
		average := float64(v.score) / float64(v.count)
		if maxAverage < int(average) {
			maxAverage = int(average)
		}
	}

	return maxAverage
}
