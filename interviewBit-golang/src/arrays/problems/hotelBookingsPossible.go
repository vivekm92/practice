package arrayProblems

import (
	"sort"
)

type Pair struct {
	t int
	a bool
}

func HotelBookings(A []int, B []int, C int) bool {

	var ps []Pair
	for _, v := range A {
		p := new(Pair)
		p.t = v
		p.a = true
		ps = append(ps, *p)
	}

	for _, v := range B {
		p := new(Pair)
		p.t = v
		p.a = false
		ps = append(ps, *p)
	}

	sort.Slice(ps, func(i, j int) bool {

		if ps[i].t == ps[j].t {
			return true
		}

		return ps[i].t < ps[j].t
	})

	k := 0
	for _, v := range ps {
		if v.a {
			k++
		} else {
			k--
		}

		if k > C {
			return false
		}
	}

	return true
}
