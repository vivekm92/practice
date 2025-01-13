package stringProblems

/*
	problem : https://www.interviewbit.com/problems/zigzag-string/
*/

// T(n) : O(n), S(n) : O(n)
func ZigZagString(A string, B int) string {

	n := len(A)
	if B == 1 || B >= n {
		return A
	}

	idx, step := 0, 1
	t := make([][]rune, B)
	for _, v := range A {
		t[idx] = append(t[idx], v)
		if idx == 0 {
			step = 1
		} else if idx == B-1 {
			step = -1
		}
		idx += step
	}

	res := ""
	for _, v1 := range t {
		for _, v2 := range v1 {
			res += string(v2)
		}
	}

	return res
}

// T(n) : O(n), S(n) : O(n)
func ZigZagString1(A string, B int) string {

	n := len(A)
	if B <= 1 || B >= n {
		return A
	}

	idx, step := 0, 1
	t := make([]string, B)
	for _, v := range A {
		if idx == 0 {
			step = 1
		} else if idx == B-1 {
			step = -1
		}
		t[idx] += string(v)
		idx += step
	}

	var res string
	for _, v := range t {
		res += v
	}

	return res
}
