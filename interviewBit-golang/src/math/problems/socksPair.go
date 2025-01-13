package mathProblems

/*
	Problem : https://www.interviewbit.com/problems/socks-pair/
*/

// T(n) : O(n) , S(n) : O(n)
func SocksPairs(A []int) int {

	m := make(map[int]int, 0)
	for _, v := range A {
		m[v]++
	}

	count := 0
	for _, v := range m {
		count += v / 2
	}

	return count
}
