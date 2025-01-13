package stringProblems

/*
	Problem : https://www.interviewbit.com/problems/valid-ip-addresses/
*/

// T(n) : O(n3) ; S(n) :O(1)
func RestoreIpAddresses(A string) []string {

	n := len(A)
	res := make([]string, 0)
	for i := 0; i < n; i++ {
		if !IsValidIPDigit(A, 0, i) {
			continue
		}

		for j := i + 1; j < n; j++ {
			if !IsValidIPDigit(A, i+1, j) {
				continue
			}

			for k := j + 1; k < n; k++ {
				if IsValidIPDigit(A, j+1, k) && IsValidIPDigit(A, k+1, n-1) {
					t := A[:i+1] + "." + A[i+1:j+1] + "." + A[j+1:k+1] + "." + A[k+1:]
					res = append(res, t)
				}
			}
		}
	}

	return res
}

func IsValidIPDigit(A string, start, end int) bool {
	if start >= len(A) || end >= len(A) {
		return false
	}
	if A[start] == '0' {
		return start == end
	}
	var t int
	for i := start; i <= end; i++ {
		t = t*10 + int(A[i]-'0')
		if t < 0 || t > 255 {
			return false
		}
	}
	return true
}
