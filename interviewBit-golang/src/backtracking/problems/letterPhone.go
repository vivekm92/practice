package backtracking

/*
	Problem : https://www.interviewbit.com/problems/letter-phone/
*/

func solveLetterPhone(A string, B string, C map[rune]string, D *[]string, idx int) {

	if idx == len(A) {
		*D = append(*D, B)
		return
	}

	s := C[rune(A[idx])]
	for i := 0; i < len(s); i++ {
		B += string(s[i])
		solveLetterPhone(A, B, C, D, idx+1)
		B = B[:len(B)-1]
	}
}

// T(n) : O() ; S(n) : O()
func LetterPhone(A string) []string {

	res := make([]string, 0)
	if len(A) == 0 {
		return res
	}

	lookup := make(map[rune]string)
	lookup['0'] = "0"
	lookup['1'] = "1"
	lookup['2'] = "abc"
	lookup['3'] = "def"
	lookup['4'] = "ghi"
	lookup['5'] = "jkl"
	lookup['6'] = "mno"
	lookup['7'] = "pqrs"
	lookup['8'] = "tuv"
	lookup['9'] = "wxyz"

	solveLetterPhone(A, "", lookup, &res, 0)
	return res
}
