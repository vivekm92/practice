package stringProblems

/*
	Problem : https://www.interviewbit.com/problems/roman-to-integer/
*/

// T(n) : O(n) , S(n) : O(1)
func RomanToInteger(A string) int {

	res := 0
	i, n := 0, len(A)
	for i < n {
		if i < n-1 && A[i:i+2] == "CM" {
			res += 900
			i++
		} else if i < n-1 && A[i:i+2] == "CD" {
			res += 500
			i++
		} else if i < n-1 && A[i:i+2] == "XC" {
			res += 90
			i++
		} else if i < n-1 && A[i:i+2] == "XL" {
			res += 50
			i++
		} else if i < n-1 && A[i:i+2] == "IX" {
			res += 9
			i++
		} else if i < n-1 && A[i:i+2] == "IV" {
			res += 4
			i++
		} else if string(A[i]) == "M" {
			res += 1000
		} else if string(A[i]) == "D" {
			res += 500
		} else if string(A[i]) == "C" {
			res += 100
		} else if string(A[i]) == "L" {
			res += 50
		} else if string(A[i]) == "X" {
			res += 10
		} else if string(A[i]) == "V" {
			res += 5
		} else {
			res += 1
		}
		i++
	}

	return res
}
