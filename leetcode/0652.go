package leetcode

import "fmt"

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	var treeMp = make(map[string]*TreeNode)
	var treeRepeatCountMp = make(map[string]int8)

	traverseBuildMp(root, treeMp, treeRepeatCountMp)

	res := make([]*TreeNode, 0, len(treeMp))
	for _, v := range treeMp {
		res = append(res, v)
	}
	return res
}

func traverseBuildMp(root *TreeNode, treeMp map[string]*TreeNode, treeRepeatCountMp map[string]int8) {
	if root == nil {
		return
	}

	traverseBuildMp(root.Left, treeMp, treeRepeatCountMp)
	traverseBuildMp(root.Right, treeMp, treeRepeatCountMp)

	// 序列化
	serializeStr := customSerialize(root)
	if _, ok := treeRepeatCountMp[serializeStr]; ok {
		treeMp[serializeStr] = root
	} else {
		treeRepeatCountMp[serializeStr] = 0
	}
}

func customSerialize(node *TreeNode) string {
	if node == nil {
		return "#"
	}

	sl := customSerialize(node.Left)
	sr := customSerialize(node.Right)
	return fmt.Sprintf("%s,%s,%d", sl, sr, node.Val)
}
