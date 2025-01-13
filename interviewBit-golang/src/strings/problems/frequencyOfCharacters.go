package stringProblems

/*
	Problem : https://www.interviewbit.com/problems/frequency-of-characters/
*/

// T(n) : O(n) , S(n) : O(1)
func FrequencyOfCharacters(A string) []int {

	res := make([]int, 26)
	for _, v := range A {
		res[v-'a']++
	}

	return res
}
