package mathProblems

/*
	Problem : https://www.interviewbit.com/problems/round-table/
*/

// T(n) : O(n), S(n) : O(1)
func RoundTable(A int) int {

	// result == 2! * (A-1)!
	// total number of people = A + 1
	// 2 persons can be arranged in 2! ways
	// considering Ram & Shyam as one entity
	// A + 1 number of people can be arranged in circle in (A + 1 - 1 - 1)! ways
	// Total = 2! * (A-1)! ways

	res := 1
	const MOD = 1000000007
	for i := 1; i < A; i++ {
		res = (res % MOD * i % MOD) % MOD
	}

	return (res % MOD * 2 % MOD) % MOD
}
