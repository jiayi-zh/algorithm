package leetcode

func isValidBST(root *TreeNode) bool {
	return subIsValidBST(root, nil, nil)
}

func subIsValidBST(root *TreeNode, min, max *TreeNode) bool {
	if root == nil {
		return true
	}
	if min != nil && root.Val <= min.Val {
		return false
	}
	if max != nil && root.Val >= max.Val {
		return false
	}
	return subIsValidBST(root.Left, min, root) && subIsValidBST(root.Right, root, max)
}
