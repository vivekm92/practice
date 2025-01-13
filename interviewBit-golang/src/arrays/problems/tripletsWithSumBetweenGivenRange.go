package arrayProblems

import (
	"sort"
	"strconv"
)

/*
	Problem : https://www.interviewbit.com/problems/triplets-with-sum-between-given-range/
*/

// T(n) : O(nlogn), S(n) : O(n)
func TripletsExists(A []string) int {

	var A1 []float64
	for _, v := range A {
		val, _ := strconv.ParseFloat(v, 64)
		A1 = append(A1, val)
	}

	sort.Float64s(A1)

	var a, b, c []float64
	for _, v := range A1 {
		if v > 0.0 && v < 2.0/3.0 {
			a = append(a, v)
		} else if v >= 2.0/3.0 && v <= 1.0 {
			b = append(b, v)
		} else if v > 1.0 && v < 2.0 {
			c = append(c, v)
		}
	}

	a1, b1, c1 := len(a), len(b), len(c)
	if a1 >= 3 {
		if a[a1-1]+a[a1-2]+a[a1-3] > 1 && a[a1-1]+a[a1-2]+a[a1-3] < 2 {
			return 1
		}
	}

	if a1 >= 2 && b1 >= 1 {
		if a[0]+a[1]+b[b1-1] > 1 && a[0]+a[1]+b[b1-1] < 2 {
			return 1
		} else if b[0]+a[a1-1]+a[a1-2] > 1 && b[0]+a[a1-1]+a[a1-2] < 2 {
			return 1
		}
	}

	if a1 >= 2 && c1 >= 1 {
		if c[0]+a[0]+a[1] > 1 && c[0]+a[0]+a[1] < 2 {
			return 1
		}
	}

	if a1 >= 1 && b1 >= 2 {
		if a[0]+b[0]+b[1] > 1 && a[0]+b[0]+b[1] < 2 {
			return 1
		}
	}

	if a1 >= 1 && b1 >= 1 && c1 >= 1 {
		if a[0]+b[0]+c[0] > 1 && a[0]+b[0]+c[0] < 2 {
			return 1
		}
	}

	return 0
}
