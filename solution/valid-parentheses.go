package solution

import "container/list"

func isValid(s string) bool {
	parenthesesMatch := map[int32]int32{
		'(': ')',
		'[': ']',
		'{': '}'}
	stack := list.New()
	for _, v := range s {
		first := stack.Front()
		if first == nil || parenthesesMatch[first.Value.(int32)] != v {
			stack.PushFront(v)
			continue
		}
		stack.Remove(first)
	}
	return stack.Len() == 0
}
