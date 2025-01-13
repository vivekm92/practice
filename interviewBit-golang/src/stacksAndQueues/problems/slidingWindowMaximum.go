package stacksAndQueues

import (
	"container/heap"
)

/*
	Problem : https://www.interviewbit.com/problems/sliding-window-maximum/
*/

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	val := old[n-1]
	*h = old[:n-1]
	return val
}

// T(n) : O(n) ; S(n) : O()n
func SlidingMaximum(A []int, B int) []int {

	var pq IntHeap
	for i := 0; i < B; i++ {
		pq = append(pq, A[i])
	}
	heap.Init(&pq)

	n, lookup := len(A), make(map[int]int, 0)
	res := make([]int, 0)
	res = append(res, pq[0])
	lookup[A[0]] = 1
	for i := B; i < n; i++ {
		for x, ok := lookup[pq[0]]; ok && x > 0 && pq.Len() > 0; {
			lookup[pq[0]]--
			heap.Pop(&pq)
			if pq.Len() > 0 {
				x, ok = lookup[pq[0]]
			}
		}

		heap.Push(&pq, A[i])
		res = append(res, pq[0])
		lookup[A[i-B+1]]++
	}

	return res
}
