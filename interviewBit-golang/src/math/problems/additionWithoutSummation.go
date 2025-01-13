package mathProblems

func AddBitwise(A, B int) int {

	for B != 0 {
		var carry uint = uint(A & B)
		A = A ^ B
		B = int(carry << 1)
	}

	return A
}
