package arrayProblems

func RepeatAndMissingNumber(A []int) []int {

	n := uint64(len(A))
	s_n := (n * (n + 1)) / 2
	s_n2 := (n * (n + 1) * (2*n + 1)) / 6

	for _, v := range A {
		s_n -= uint64(v)
		s_n2 -= uint64(v) * uint64(v)
	}

	missing := (s_n + s_n2/s_n) / 2
	repeat := -1*int(s_n) + int(missing)

	return []int{int(repeat), int(missing)}
}
