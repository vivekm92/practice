package mathProblems

/*
	Problem : https://www.interviewbit.com/problems/excel-column-title/
*/

// T(n) : O(logn), S(n) : O(1)
func ConvertToTitle(A int) string {

	lookup := make(map[int]string, 0)
	for i := 65; i <= 90; i++ {
		lookup[i-65] = string(byte(i))
	}

	res := ""
	for A > 0 {
		res += lookup[(A-1)%26]
		A = (A - 1) / 26
	}

	r := []rune(res)
	for i := 0; i < len(r)/2; i++ {
		r[i], r[len(r)-i-1] = r[len(r)-i-1], r[i]
	}

	return string(r)
}
