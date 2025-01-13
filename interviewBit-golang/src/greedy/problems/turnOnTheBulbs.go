package greedy

/*
	Problem : https://www.interviewbit.com/problems/turn-on-the-bulbs/
*/

// T(n) : O(1) ; S(n) : O(1)
func TurnOnTheBulbs(A int) int {

	if A == 1 {
		return 1
	} else if A > 1 && A <= 6 {
		return 2
	}

	t := A / 6
	u := A % 6
	if u == 0 {
		return 2 * t
	}

	return 2*t + TurnOnTheBulbs(u)
}
