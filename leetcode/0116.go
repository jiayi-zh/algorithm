package leetcode

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	connectLR(root.Left, root.Right)
	return root
}

func connectLR(left *Node, right *Node) {
	if left == nil || right == nil {
		return
	}

	left.Next = right

	connectLR(left.Left, left.Right)
	connectLR(right.Left, right.Right)
	connectLR(left.Right, right.Left)
}
