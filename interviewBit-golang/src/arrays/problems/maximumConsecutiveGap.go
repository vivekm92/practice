package arrayProblems

import (
	"math"
	"sort"
)

/*
	problem : https://www.interviewbit.com/problems/maximum-consecutive-gap/
*/

// T(n) : O(nlogn), S(n) : O(1)
func MaximumGap(A []int) int {

	sort.Slice(A, func(i int, j int) bool {
		return A[i] < A[j]
	})

	n, maxGap := len(A), 0
	for i := 0; i < n-1; i++ {
		if A[i+1]-A[i] > maxGap {
			maxGap = A[i+1] - A[i]
		}
	}

	return maxGap
}

// T(n) : O(n), S(n) : O(n)
func MaximumGap1(A []int) int {

	n := len(A)
	if n < 2 {
		return 0
	}

	maxElement, minElement := math.MinInt32, math.MaxInt32
	for _, v := range A {
		if v > maxElement {
			maxElement = v
		}
		if v < minElement {
			minElement = v
		}
	}

	gap := (maxElement-minElement-1)/(n-1) + 1
	if maxElement == minElement {
		return 0
	}
	buckets := make([]Bucket, n)
	for i := 0; i < n; i++ {
		buckets[i].MinEle = math.MaxInt32
		buckets[i].MaxEle = math.MinInt32
	}
	for i := 0; i < n; i++ {
		idx := (A[i] - minElement) / gap
		if A[i] > buckets[idx].MaxEle {
			buckets[idx].MaxEle = A[i]
		}

		if A[i] < buckets[idx].MinEle {
			buckets[idx].MinEle = A[i]
		}

	}

	maxGap := 0
	for i := 0; i < n-1; i++ {
		v2 := buckets[i+1].MinEle
		v1 := buckets[i].MaxEle

		if v2-v1 > maxGap && v2 != math.MaxInt32 && v1 != math.MinInt32 {
			maxGap = v2 - v1
		}
	}

	return maxGap
}

type Bucket struct {
	MinEle int
	MaxEle int
}
