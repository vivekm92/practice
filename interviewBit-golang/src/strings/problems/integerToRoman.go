package stringProblems

/*
	Problem : https://www.interviewbit.com/problems/integer-to-roman/
*/

// T(n) : O(logn) , S(n) : O(1)
func IntegerToRoman(A int) string {

	var res string
	for A > 0 {
		if A == 4 {
			res += "IV"
			A -= 4
		} else if A == 9 {
			res += "IX"
			A -= 9
		} else if A >= 40 && A < 50 {
			res += "XL"
			A -= 40
		} else if A >= 90 && A < 100 {
			res += "XC"
			A -= 90
		} else if A >= 400 && A < 500 {
			res += "CD"
			A -= 400
		} else if A >= 900 && A < 1000 {
			res += "CM"
			A -= 900
		} else if A >= 1000 {
			res += "M"
			A -= 1000
		} else if A >= 500 {
			res += "D"
			A -= 500
		} else if A >= 100 {
			res += "C"
			A -= 100
		} else if A >= 50 {
			res += "L"
			A -= 50
		} else if A >= 10 {
			res += "X"
			A -= 10
		} else if A >= 5 {
			res += "V"
			A -= 5
		} else {
			res += "I"
			A -= 1
		}
	}

	return res
}
