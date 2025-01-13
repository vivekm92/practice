package mathProblems

/*
	Problem : https://www.interviewbit.com/problems/distribute-in-circle/
*/

// T(n) : O(1), S(1) : O(1)
func DistributeInCircle(A, B, C int) int {
	return (C + A - 1) % B
}
