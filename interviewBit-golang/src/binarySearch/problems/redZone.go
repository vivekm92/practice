package bsearch

import "math"

/*
	Problem : https://www.interviewbit.com/problems/red-zone/
*/

type Point struct {
	X float64
	Y float64
}

func (p *Point) Abs() float64 {
	x, y := p.X, p.Y
	xd, yd := float64(x), float64(y)
	return math.Sqrt(xd*xd + yd*yd)
}

func checkForRedZone(A [][]int, B int, r int) bool {

	n := len(A)
	arr := make([]Point, 0)
	for i := 0; i < n; i++ {
		var p Point
		p.X, p.Y = float64(A[i][0]), float64(A[i][1])
		arr = append(arr, p)
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			var g Point
			g.X = arr[j].X - arr[i].X
			g.Y = arr[j].Y - arr[i].Y
			d := g.Abs()

			rf := float64(r)
			if d > 2*rf {
				continue
			}

			var m Point
			m.X = (arr[j].X + arr[j].X) / 2.0
			m.Y = (arr[j].Y + arr[j].Y) / 2.0

			h := math.Sqrt(rf*rf - d*d/4.0)

			var p Point
			p.X = -1 * g.Y * (h / d)
			p.Y = g.X * (h / d)
			c1, c2 := 2, 2
			for l := 0; l < n; l++ {
				var v1, v2 Point
				v1.X, v1.Y = arr[l].X-(m.X-p.X), arr[l].Y-(m.Y-p.Y)
				v2.X, v2.Y = arr[l].X-(m.X+p.X), arr[l].Y-(m.Y-p.Y)
				if l == i || l == j {
					continue
				} else if v1.Abs() <= rf {
					c1++
				} else if v2.Abs() <= rf {
					c2++
				}
			}

			// if c >= c1 && c >= c2 {
			// 	c = c
			// } else if c >= c1 && c < c2 {
			// 	return c2
			// } else {

			// }

		}
	}
	return true
}

// T(n) :
func RedZone(A [][]int, B int) int {

	l, r := 0, 100001
	firstDay := -1
	for l <= r {
		mid := l + (r-l)/2
		if checkForRedZone(A, B, mid) {
			firstDay = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return firstDay
}
