package leetcode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := new(ListNode)
	bfo := dummy
	var v1, v2, nv int
	bitFlag := false
	for {
		if l1 == nil && l2 == nil && !bitFlag {
			break
		}
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		} else {
			v1 = 0
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		} else {
			v2 = 0
		}

		nv = v1 + v2
		if bitFlag {
			nv += 1
		}
		if nv >= 10 {
			nv = nv - 10
			bitFlag = true
		} else {
			bitFlag = false
		}

		cur := new(ListNode)
		cur.Val = nv
		bfo.Next = cur
		bfo = cur
	}
	return dummy.Next
}
