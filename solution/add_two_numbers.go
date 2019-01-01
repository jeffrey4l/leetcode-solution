package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var root *ListNode
	var tmp *ListNode
	carry := false

	for l1 != nil || l2 != nil {
		a1, a2 := 0, 0
		if l1 != nil {
			a1 = l1.Val
		}
		if l2 != nil {
			a2 = l2.Val
		}
		dataVal := a1 + a2
		if carry {
			carry = false
			dataVal += 1
		}
		if dataVal >= 10 {
			dataVal = dataVal - 10
			carry = true
		}
		data := &ListNode{Val: dataVal, Next: nil}
		if tmp == nil {
			tmp = data
			root = data
		} else {
			tmp.Next = data
			tmp = data
		}

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}

	}

	if carry {
		tmp.Next = &ListNode{Val: 1, Next: nil}
	}
	return root
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {

	var root = &ListNode{}
	tmp := root
	carry := 0

	for l1 != nil || l2 != nil {
		a1, a2 := 0, 0
		if l1 != nil {
			a1 = l1.Val
		}
		if l2 != nil {
			a2 = l2.Val
		}
		dataVal := a1 + a2 + carry
		carry = dataVal / 10
		data := &ListNode{Val: dataVal % 10, Next: nil}

		tmp.Next = data
		tmp = data

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}

	}

	if carry != 0 {
		tmp.Next = &ListNode{Val: 1, Next: nil}
	}
	return root.Next
}
