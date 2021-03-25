package leetcode

func buildTree1(preorder []int, inorder []int) *TreeNode {
	return reductionTree1(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1)
}

func reductionTree1(preorder []int, pl, pr int, inorder []int, il, ir int) *TreeNode {
	if pl > pr || il > ir {
		return nil
	}
	root := new(TreeNode)
	root.Val = preorder[pl]

	iri := findValIndex(inorder, root.Val, il, ir)

	// 不管是前序，中序，后序， 根节点两边的个数都是一致的
	root.Left = reductionTree1(preorder, pl+1, pl+(iri-il), inorder, il, iri-1)
	root.Right = reductionTree1(preorder, pl+(iri-il)+1, pr, inorder, iri+1, ir)

	return root
}
