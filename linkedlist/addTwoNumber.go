package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{-1, nil}
	head := res

	var carry, sum int

	for l1 != nil && l2 != nil {
		i, j := l1.Val, l2.Val

		total := carry + i + j
		carry = total / 10
		sum = total % 10

		res.Next = &ListNode{sum, nil}
		res = res.Next

		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		total := carry + l1.Val

		carry = total / 10
		sum = total % 10

		res.Next = &ListNode{sum, nil}
		res = res.Next
		l1 = l1.Next
	}

	for l2 != nil {
		total := carry + l2.Val

		carry = total / 10
		sum = total % 10

		res.Next = &ListNode{sum, nil}
		res = res.Next
		l2 = l2.Next
	}

	if carry > 0 {
		res.Next = &ListNode{carry, nil}
	}

	return head.Next
}

func addTwoNumbersDriver() {
	var l1, l2 *ListNode

	l1 = l1.add(9)
	l1 = l1.add(9)
	l1 = l1.add(9)
	l1 = l1.add(9)
	l1 = l1.add(9)
	l1 = l1.add(9)
	l1 = l1.add(9)

	l2 = l2.add(9)
	l2 = l2.add(9)
	l2 = l2.add(9)
	l2 = l2.add(9)

	res := addTwoNumbers(l1, l2)
	res.display()
}
