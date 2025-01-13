package arrayProblems

/*
  Problem : https://www.interviewbit.com/problems/array-sum/

  Solution :

  1) Two Pointer Approach

	- Keep pointers at the end of each array.
	- keep adding the values pointed at each array and decrease the pointer value.
	- reverse the resultant array.

  Follow up :

  1) Try to do the above problem without using any additional space.

	- Use the longer input array as the resultant array.
	- Save the sum in resultant array using two pointer approach.
	- if carry is present then append the resultant array with carry.
	- right shift all the elemnets by 1.
	- set carry as the first element of the array.
	- return the array.

*/

// T(n) : O(n), S(n) : O(n)
func AddArrays(A []int, B []int) []int {

	arr := make([]int, 0)
	iA, iB, carry := len(A)-1, len(B)-1, 0
	for iA >= 0 || iB >= 0 {
		valA, valB := 0, 0
		if iA >= 0 {
			valA = A[iA]
			iA--
		}
		if iB >= 0 {
			valB = B[iB]
			iB--
		}

		sum := valA + valB + carry
		arr = append(arr, sum%10)
		carry = sum / 10
	}

	if carry != 0 {
		arr = append(arr, carry)
	}

	// reverse the resultant array.
	n := len(arr)
	for i := 0; i < int(n/2); i++ {
		arr[i], arr[n-1-i] = arr[n-1-i], arr[i]
	}

	return arr
}

// T(n) : O(n), S(n) : O(1)
func AddArraysFollowUp(A []int, B []int) []int {

	nA, nB := len(A)-1, len(B)-1
	res, iC := A, nA
	if nB > nA {
		res = B
		iC = nB
	}

	iA, iB, carry := nA, nB, 0
	for iA >= 0 || iB >= 0 {
		valA, valB := 0, 0
		if iA >= 0 {
			valA = A[iA]
			iA--
		}

		if iB >= 0 {
			valB = B[iB]
			iB--
		}

		sum := valA + valB + carry
		res[iC] = sum % 10
		iC--
		carry = sum / 10
	}

	if carry != 0 {
		return append([]int{carry}, res...)
	}

	return res
}

// T(n) : O(n), S(n) : O(1)
func AddArraysFollowUp1(A []int, B []int) []int {

	nA, nB := len(A)-1, len(B)-1
	res, iC := A, nA
	if nB > nA {
		res = B
		iC = nB
	}

	iA, iB, carry := nA, nB, 0
	for iA >= 0 || iB >= 0 {
		valA, valB := 0, 0
		if iA >= 0 {
			valA = A[iA]
			iA--
		}

		if iB >= 0 {
			valB = B[iB]
			iB--
		}

		sum := valA + valB + carry
		res[iC] = sum % 10
		iC--
		carry = sum / 10
	}

	if carry != 0 {
		res = append(res, carry)
		nC := len(res)
		for i := nC - 1; i > 0; i-- {
			res[i] = res[i-1]
		}
		res[0] = carry
	}

	return res
}
