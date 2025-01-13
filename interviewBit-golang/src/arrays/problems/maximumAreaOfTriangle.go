package arrayProblems

import (
	"math"
)

/*
	Problem : https://www.interviewbit.com/problems/maximum-area-of-triangle/
*/

// T(n) : O(n*m), S(n) : O(n)
func MaxAreaOfTriangle(A []string) int {

	m, n := len(A), len(A[0])
	rs, re := make([]int, n), make([]int, n)
	gs, ge := make([]int, n), make([]int, n)
	bs, be := make([]int, n), make([]int, n)

	for i := 0; i < n; i++ {
		rs[i] = -1
		re[i] = -1
		gs[i] = -1
		ge[i] = -1
		bs[i] = -1
		be[i] = -1
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if rs[i] == -1 && A[j][i] == 'r' {
				rs[i] = j
			}
			if A[j][i] == 'r' {
				re[i] = j
			}

			if gs[i] == -1 && A[j][i] == 'g' {
				gs[i] = j
			}
			if A[j][i] == 'g' {
				ge[i] = j
			}

			if bs[i] == -1 && A[j][i] == 'b' {
				bs[i] = j
			}
			if A[j][i] == 'b' {
				be[i] = j
			}
		}
	}

	maxArea := 0.0
	for i := 0; i < n; i++ {
		bmax, rp, gp, bp := findMaxB(rs[i], re[i], gs[i], ge[i], bs[i], be[i])
		if bmax == -1 {
			continue
		}

		var t []int
		if !rp && !gp && !bp {
			continue
		} else if !rp {
			t = re
		} else if !gp {
			t = ge
		} else if !bp {
			t = be
		}

		for j := 0; j < n; j++ {
			if t[j] == -1 {
				continue
			}
			bf := float64(bmax)
			hf := float64(absDiff(j, i))
			area := 0.5 * bf * hf
			if area > maxArea {
				maxArea = area
			}
		}
	}

	return int(math.Ceil(maxArea))
}

func findMaxB(rs, re, gs, ge, bs, be int) (int, bool, bool, bool) {

	if re == -1 && ge == -1 && be == -1 {
		return -1, false, false, false
	} else if (re == -1 && ge == -1) || (ge == -1 && be == -1) || (be == -1 && re == -1) {
		return -1, false, false, false
	} else if re == -1 {
		if absDiff(bs, ge) > absDiff(gs, be) {
			return absDiff(bs, ge), false, true, true
		} else {
			return absDiff(gs, be), false, true, true
		}
	} else if ge == -1 {
		if absDiff(rs, be) > absDiff(bs, re) {
			return absDiff(rs, be), true, false, true
		} else {
			return absDiff(bs, re), true, false, true
		}
	} else if be == -1 {
		if absDiff(rs, ge) > absDiff(gs, re) {
			return absDiff(rs, ge), true, true, false
		} else {
			return absDiff(gs, re), true, true, false
		}
	} else {
		a1 := absDiff(rs, ge)
		if a1 < absDiff(gs, re) {
			a1 = absDiff(gs, re)
		}
		a2 := absDiff(gs, be)
		if a2 < absDiff(ge, bs) {
			a2 = absDiff(ge, bs)
		}
		a3 := absDiff(bs, re)
		if a3 < absDiff(rs, be) {
			a3 = absDiff(rs, be)
		}

		if a1 >= a2 && a1 >= a3 {
			return a1, true, true, false
		} else if a2 >= a1 && a2 >= a3 {
			return a2, false, true, true
		} else {
			return a3, true, false, true
		}
	}

}

func absDiff(a, b int) int {
	if a > b {
		return a - b + 1
	}
	return b - a + 1
}
