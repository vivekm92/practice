package stringProblems

/*
	Problem : https://www.interviewbit.com/problems/pretty-json/
*/

// T(n) : O(n) . S(n) : O(1)
func PrettyJson(A string) []string {

	var t string
	indentCount, res, n := 0, make([]string, 0), len(A)
	for i, v := range A {
		if v == ' ' {
			continue
		} else if v == '{' || v == '[' {
			t = addIndent(indentCount) + string(v)
			res = append(res, t)
			indentCount++
			t = addIndent(indentCount)
		} else if v == '}' || v == ']' {
			if len(t) > 0 && t[len(t)-1] != '\t' {
				res = append(res, t)
			}
			if i+1 < n && A[i+1] == ',' {
				indentCount--
				t = addIndent(indentCount) + string(v)
				continue
			}
			indentCount--
			t = addIndent(indentCount) + string(v)
			res = append(res, t)
			t = addIndent(indentCount)
		} else if v == ',' {
			t += string(v)
			res = append(res, t)
			t = addIndent(indentCount)
		} else if v == ':' {
			if i+1 < n && (A[i+1] == '{' || A[i+1] == '[') {
				t += string(v)
				res = append(res, t)
				t = addIndent(indentCount)
			} else {
				t += string(v)
			}
		} else {
			t += string(v)
		}
	}
	return res
}

func addIndent(indentCount int) string {
	var res string
	for indentCount > 0 {
		res += string('\t')
		indentCount--
	}
	return res
}
