package leetcode

func constructMaximumBinaryTree(nums []int) *TreeNode {
	return constructBinaryTree(nums, 0, len(nums)-1)
}

func constructBinaryTree(nums []int, l, r int) *TreeNode {
	if l > r {
		return nil
	}
	maxIndex := getMax(nums, l, r)
	root := &TreeNode{
		Val: nums[maxIndex],
	}

	root.Left = constructBinaryTree(nums, l, maxIndex-1)
	root.Right = constructBinaryTree(nums, maxIndex+1, r)
	return root
}

func getMax(nums []int, l, r int) int {
	max := -1
	maxIndex := -1
	for ; l <= r; l++ {
		if nums[l] > max {
			max = nums[l]
			maxIndex = l
		}
	}
	return maxIndex
}
