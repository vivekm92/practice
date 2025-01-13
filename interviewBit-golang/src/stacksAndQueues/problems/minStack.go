package stacksAndQueues

/*
	Problem : https://www.interviewbit.com/problems/min-stack/
*/

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

type Pair struct {
	First, Second int
}

type MinStack struct {
	Ele []Pair
}

func MinStackConstructor() MinStack {
	return MinStack{}
}

// T(n) : O(1)
func (m *MinStack) Push(x int) {
	var p Pair
	if len(m.Ele) == 0 {
		p.First, p.Second = x, x
	} else {
		p.First = x
		if x < m.Ele[len(m.Ele)-1].Second {
			p.Second = x
		} else {
			p.Second = m.Ele[len(m.Ele)-1].Second
		}
	}
	m.Ele = append(m.Ele, p)
}

// T(n) : O(1)
func (m *MinStack) Pop() {
	if len(m.Ele) == 0 {
		return
	}
	m.Ele = m.Ele[:len(m.Ele)-1]
}

// T(n) : O(1)
func (m *MinStack) Top() int {
	if len(m.Ele) == 0 {
		return -1
	}
	return m.Ele[len(m.Ele)-1].First
}

// T(n) : O(1)
func (m *MinStack) GetMin() int {
	if len(m.Ele) == 0 {
		return -1
	}
	return m.Ele[len(m.Ele)-1].Second
}
