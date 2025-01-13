package backtrackingExamples

func solveMod(A, B, C int) int {
	if B == 0 {
		return 1
	} else if B == 1 {
		return A
	}

	res := solveMod(A, B/2, C)
	t := ((res % C) * (res % C)) % C
	if B%2 == 0 {
		return t
	}

	return ((A % C) * (t % C)) % C
}

// T(n): O(logN), S(n) : O(logN)
func Mod(A int, B int, C int) int {
	if A == 0 {
		return 0
	}
	for A < 0 {
		A += C
	}
	return solveMod(A, B, C)
}
