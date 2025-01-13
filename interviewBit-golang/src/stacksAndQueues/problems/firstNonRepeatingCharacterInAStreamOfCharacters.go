package stacksAndQueues

import (
	"container/list"
)

/*
	problem : https://www.interviewbit.com/problems/first-non-repeating-character-in-a-stream-of-characters/
*/

// T(n) : O(n), S(n) : O(n)
func NonRepeatingCharacters(A string) string {

	var res string
	q := list.New()
	lookup := make(map[rune]int)
	for _, ch := range A {
		if _, ok := lookup[ch]; !ok {
			q.PushBack(ch)
			lookup[ch] = 1
		} else {
			lookup[ch] = -1
		}

		if q.Len() > 0 {
			front := q.Front().Value.(rune)
			for q.Len() > 0 && lookup[front] == -1 {
				q.Remove(q.Front())
				if q.Len() == 0 {
					break
				}
				front, _ = q.Front().Value.(rune)
			}
		}

		if q.Len() == 0 {
			res += "#"
		} else {
			front, _ := q.Front().Value.(rune)
			res += string(front)
		}
	}

	return res
}

// T(n) : O(n), S(n) : O(n)
func NonRepeatingCharacters1(A string) string {

	q, res := make([]byte, 0, len(A)), make([]byte, 0, len(A))
	mp := make(map[byte]int)
	for j := 0; j < len(A); j++ {
		if v, ok := mp[A[j]]; ok {
			mp[A[j]] = v + 1
		} else {
			mp[A[j]] = 1
			q = append(q, A[j])
		}

		i := 0
		for ; i < len(q); i++ {
			if mp[q[i]] <= 1 {
				res = append(res, q[i])
				break
			}
		}

		if i == len(q) {
			res = append(res, '#')
		}
		q = q[i:]
	}

	return string(res)
}
