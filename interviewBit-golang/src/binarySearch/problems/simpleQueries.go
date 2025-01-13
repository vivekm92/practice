package bsearch

import (
	"interviewBit/src/utils"
	"sort"
)

/*
	Problem : https://www.interviewbit.com/problems/simple-queries
*/

const MAX = 100010
const MOD = 1000000007

var prod []int = make([]int, MAX)

// T(n) : O(n), S(n) : O(n)
func ComputeProductOfDivisors() []int {
	if prod[1] != 0 {
		return prod
	}

	prod[1] = 1
	for i := 2; i < MAX; i++ {
		prod[i] = i
	}

	for i := 2; i < MAX; i++ {
		for j := 2 * i; j < MAX; j += i {
			prod[j] = (i * prod[j]) % MOD
		}
	}

	return prod
}

// T(n) : O(n), S(n) : O(n)
func SubArrayCount(A []int) []int {

	n := len(A)
	freq := make([]int, n)
	st := utils.NewStack[int](n)
	st.Push(0)
	for i := 1; i < n; i++ {
		topEle, ok := st.Peek()
		for !st.IsEmpty() && ok && A[topEle] <= A[i] {
			idx, _ := st.Pop()
			left := idx + 1
			right := i - idx
			if !st.IsEmpty() {
				t, _ := st.Peek()
				left = idx - t
			}
			freq[idx] = left * right
		}
		st.Push(i)
	}

	for !st.IsEmpty() {
		idx, _ := st.Pop()
		left := idx + 1
		right := n - idx
		if !st.IsEmpty() {
			t, _ := st.Peek()
			left = idx - t
		}
		freq[idx] = left * right
	}

	return freq
}

// T(n) : O(NlogN), S(n) : O(N)
func SimpleQueries(A []int, B []int) []int {

	n := len(A)
	prod := ComputeProductOfDivisors()
	freq := SubArrayCount(A)

	store := make([]Pair, 0)
	for i := 0; i < n; i++ {
		var p Pair
		p.Product = prod[A[i]]
		p.Freq = freq[i]
		store = append(store, p)
	}

	sort.Slice(store, func(i, j int) bool {
		return store[i].Product > store[j].Product
	})

	for i := 1; i < n; i++ {
		store[i].Freq += store[i-1].Freq
	}

	res := make([]int, 0)
	for i := 0; i < len(B); i++ {
		l, r := 0, n
		for l < r {
			mid := l + (r-l)/2
			if store[mid].Freq < B[i] {
				l = mid + 1
			} else {
				r = mid
			}
		}
		res = append(res, store[l].Product)
	}

	return res
}

type Pair struct {
	Product int
	Freq    int
}
