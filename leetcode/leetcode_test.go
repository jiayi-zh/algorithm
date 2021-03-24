package leetcode

import (
	"testing"
)

func Test(t *testing.T) {
	constructMaximumBinaryTree([]int{3, 2, 1, 6, 0, 5})

	//tree := buildTestTree()
	//flatten(tree)

	assert(t, "aa", "aa", longestPalindrome("aa"))
}

func assert(t *testing.T, item string, except, now interface{}) {
	if except != now {
		t.Fatalf("%s except: %v, now: %v", item, except, now)
	}
	// TODO 回顾反射, 判断类型与值是否匹配
}

func buildTestTree() *TreeNode {
	l31 := TreeNode{
		Val: 1,
	}
	l32 := TreeNode{
		Val: 3,
	}
	l33 := TreeNode{
		Val: 6,
	}
	l34 := TreeNode{
		Val: 9,
	}

	l21 := TreeNode{
		Val:   2,
		Left:  &l31,
		Right: &l32,
	}
	l22 := TreeNode{
		Val:   7,
		Left:  &l33,
		Right: &l34,
	}

	root := TreeNode{
		Val:   4,
		Left:  &l21,
		Right: &l22,
	}
	return &root
}
