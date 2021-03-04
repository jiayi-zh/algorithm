package leetcode

import (
	"testing"
)

func Test(t *testing.T) {
	assert(t, "aa", "aa", longestPalindrome("aa"))
}

func assert(t *testing.T, item string, except, now interface{}) {
	if except != now {
		t.Fatalf("%s except: %v, now: %v", item, except, now)
	}
	// TODO 回顾反射, 判断类型与值是否匹配
}
