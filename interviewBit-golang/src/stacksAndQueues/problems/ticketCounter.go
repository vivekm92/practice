package stacksAndQueues

/*
	Problem : https://www.interviewbit.com/problems/ticket-counter/
*/

// T(n) : O(n) , S(n) : O(1)
func TicketCounter(A []int, B []int) int {

	n := len(A)
	count, temp := 0, 0
	for i := 0; i < n; i++ {
		k := A[i] - temp
		if k < 0 {
			count++
		} else {
			temp = temp + B[i]
		}
	}

	return count
}
