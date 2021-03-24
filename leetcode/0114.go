package leetcode

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	flatten(root.Left)
	flatten(root.Right)
	// 执行到这里的时候左右子树已经被拉成一条链表

	right := root.Right
	root.Right = root.Left
	root.Left = nil

	// 将右子树拼接到左子树后面
	p := root
	for p.Right != nil {
		p = p.Right
	}
	p.Right = right
}
