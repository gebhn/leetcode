package medium

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	getDigit := func(i int) int { return i % 10 }
	getCarry := func(i int) int { return i / 10 }

	result := &ListNode{}
	current := result

	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = getCarry(sum)
		current.Next = &ListNode{Val: getDigit(sum)}
		current = current.Next
	}

	return result.Next
}
