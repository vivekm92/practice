package stacksAndQueues

import (
	"container/list"
	"math"
)

/*
	problem : https://www.interviewbit.com/problems/hotel-service/
*/

// T(n) : O(n), S(n) : O(n)
func NearestHotel(A [][]int, B [][]int) []int {

	n, m := len(A), len(A[0])
	q := list.New()
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if A[i][j] == 1 {
				A[i][j] = 0
				q.PushBack([]int{i, j})
			} else {
				A[i][j] = math.MaxInt32
			}
		}
	}

	rows := [4]int{0, 0, 1, -1}
	cols := [4]int{1, -1, 0, 0}
	for q.Len() > 0 {
		front, _ := q.Front().Value.([]int)
		q.Remove(q.Front())
		x, y := front[0], front[1]
		for i := 0; i < 4; i++ {
			nx, ny := x+rows[i], y+cols[i]
			if nx >= 0 && nx < n && ny >= 0 && ny < m && A[nx][ny] == math.MaxInt32 {
				A[nx][ny] = A[x][y] + 1
				q.PushBack([]int{nx, ny})
			}
		}
	}

	res := make([]int, 0)
	for i := 0; i < len(B); i++ {
		x, y := B[i][0]-1, B[i][1]-1
		res = append(res, A[x][y])
	}

	return res
}
