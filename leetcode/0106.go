package leetcode

func buildTree(inorder []int, postorder []int) *TreeNode {
	return reductionTree(inorder, 0, len(inorder)-1, postorder, 0, len(postorder)-1)
}

func reductionTree(inorder []int, il, ir int, postorder []int, pl, pr int) *TreeNode {
	if il > ir || pl > pr {
		return nil
	}

	root := new(TreeNode)
	root.Val = postorder[pr]

	iri := findValIndex(inorder, root.Val, il, ir)

	root.Left = reductionTree(inorder, il, iri-1, postorder, pl, pr-(ir-iri)-1)
	root.Right = reductionTree(inorder, iri+1, ir, postorder, pr-(ir-iri), pr-1)
	return root
}

func findValIndex(arr []int, val int, si, ei int) int {
	for ; si <= ei; si++ {
		if arr[si] == val {
			return si
		}
	}
	return -1
}
