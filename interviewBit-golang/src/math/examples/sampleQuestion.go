package mathExamples

func SampleQuestion(A int, B int) int {
	const mod = 10000000
	return ((A % mod) + (B % mod)) % mod
}
