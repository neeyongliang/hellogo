package add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	var jinwei int = 0
	sumV := l1.Val + l2.Val + jinwei
	head := &ListNode{sumV % 10, nil}
	jinwei = sumV / 10
	tmp := head

	val1 := 0
	val2 := 0
	l1 = l1.Next
	l2 = l2.Next

	for l1 != nil || l2 != nil || jinwei != 0 {
		if l1 == nil {
			val1 = 0
		} else {
			val1 = l1.Val
			l1 = l1.Next
		}

		if l2 == nil {
			val2 = 0
		} else {
			val2 = l2.Val
			l2 = l2.Next
		}

		sumV = val1 + val2 + jinwei
		jinwei = sumV / 10
		sumV = sumV % 10
		tmp.Next = &ListNode{sumV, nil}
		tmp = tmp.Next
	}

	return head
}
