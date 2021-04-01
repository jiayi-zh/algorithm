package leetcode

var arr = make([]int, 0, 0)

func kthSmallest(root *TreeNode, k int) int {

	subKthSmallest(root)
	if k < len(arr) {
		return arr[k-1]
	} else {
		return 0
	}
}

func subKthSmallest(root *TreeNode) {
	if root == nil {
		return
	}
	subKthSmallest(root.Left)

	arr = append(arr, root.Val)

	subKthSmallest(root.Right)
}
