package leetcode

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	invertTree(root.Left)
	invertTree(root.Right)

	left := root.Left
	root.Left = root.Right
	root.Right = left
	return root
}
