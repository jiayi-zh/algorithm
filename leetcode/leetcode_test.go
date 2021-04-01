package leetcode

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {

	//tree := buildTestTree()

	tree := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
		},
		Right: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	res := isValidBST(tree)
	fmt.Println(res)

	//tree1 := &TreeNode{
	//	Val: 0,
	//	Left: &TreeNode{
	//		Val: 0,
	//	},
	//}

	//}
	//
	//fmt.Println(customSerialize(tree1))
	//fmt.Println(customSerialize(tree2))
	////flatten(tree)
	//
	//assert(t, "aa", "aa", longestPalindrome("aa"))
}

func assert(t *testing.T, item string, except, now interface{}) {
	if except != now {
		t.Fatalf("%s except: %v, now: %v", item, except, now)
	}
	// TODO 回顾反射, 判断类型与值是否匹配
}

func buildLeetcodeTreeTestDemo(arr []int) {

}

//func buildLeetcodeTreeTestDemoChild(arr []int, l, r int, pns []*TreeNode) {
//	cPns := make([]*TreeNode, 0, len(pns)<<1)
//	for i, pn := range pns {
//		if arr[i*2] == nil {
//
//		}
//		pn.Left = *TreeNode{
//			Val:,
//		}
//	}
//}

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
