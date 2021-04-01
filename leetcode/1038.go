package leetcode

func bstToGst(root *TreeNode) *TreeNode {
	var sum = 0

	var subBstToGst func(node *TreeNode)
	subBstToGst = func(node *TreeNode) {
		if node == nil {
			return
		}
		subBstToGst(node.Right)

		sum += node.Val
		node.Val = sum

		subBstToGst(node.Left)
	}

	subBstToGst(root)
	return root
}
