package leetcode

// TODO 待优化 -> 中间有一些重复的字串判断, 其实时没有必要的
func lengthOfLongestSubstring(s string) int {
	sRune := []rune(s)
	var max int
	for i := range sRune {
		mp := make(map[rune]int)
		var cur int
		for inner := i; inner < len(sRune); inner++ {
			if _, ok := mp[sRune[inner]]; ok {
				break
			} else {
				mp[sRune[inner]] = i
				cur++
			}
		}
		if cur > max {
			max = cur
		}
	}
	return max
}
