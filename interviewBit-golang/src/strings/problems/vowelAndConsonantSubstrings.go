package stringProblems

/*
	Prblem : https://www.interviewbit.com/problems/vowel-and-consonant-substrings/
*/

// T(n) : O(n) , S(n) : O(1)
func VowelAndConsonantSubstrings(A string) int {

	const MOD = 1000000007
	tv, tc, n := 0, 0, len(A)
	for i := 0; i < n; i++ {
		if isVowel(rune(A[i])) {
			tv++
		}
	}
	tc = n - tv
	count := 0
	for i := 0; i < n; i++ {
		if isVowel(rune(A[i])) {
			tv--
			count = (count%MOD + tc%MOD) % MOD
		} else {
			tc--
			count = (count%MOD + tv%MOD) % MOD
		}
	}
	return count
}

func isVowel(A rune) bool {
	return A == 'a' || A == 'e' || A == 'i' || A == 'o' || A == 'u'
}
