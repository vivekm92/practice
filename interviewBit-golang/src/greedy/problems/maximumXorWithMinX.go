package greedy

/*
	Problem : https://www.interviewbit.com/problems/maximum-xor-with-min-x/
*/

// T(n) : O(logn) ; S(n) : O(1)
func MaximumXorWithMinX(A int) []int {

	t1 := convertToBinary(A)

	n, t2, flag := len(t1), "", true
	for i := 0; i < n; i++ {
		if t1[i] == '0' && flag {
			flag = false
		}
		if flag {
			continue
		}

		if t1[i] == '1' {
			t2 += "0"
		} else {
			t2 += "1"
		}
	}

	var x, y int
	if A == 1 {
		x = 1
		y = 1
	} else if len(t2) == 0 {
		x = 1
		y = (A / 2) * 2
	} else {
		x = convertToDecimal(t2)
		y = convertToDecimal(t1)
	}

	return []int{x, y}
}

func convertToBinary(A int) string {

	t := ""
	for A > 0 {
		if A%2 == 0 {
			t += "0"
		} else {
			t += "1"
		}
		A /= 2
	}

	return reverse((t))
}

func reverse(A string) string {

	t, n := "", len(A)
	for i := n - 1; i >= 0; i-- {
		t += string(A[i])
	}

	return t
}

func convertToDecimal(A string) int {

	t := 0
	for _, v := range A {
		if v == '1' {
			t = 2*t + 1
		} else {
			t = 2 * t
		}
	}

	return t
}
