package leetcode

// 滑动窗口
func lengthOfLongestSubstring(s string) int {
	strRuneArr := []rune(s)

	var rightIndex, ans int
	mp := make(map[rune]int)
	for leftIndex := range strRuneArr {
		if leftIndex > 0 {
			// 这里将左边界外的元素移除，此时右边界还在重复的位置
			delete(mp, strRuneArr[leftIndex-1])
		}

		for ; rightIndex < len(strRuneArr); rightIndex++ {
			if latestRepeatIndex, exist := mp[strRuneArr[rightIndex]]; exist {
				if latestRepeatIndex >= leftIndex {
					break
				}
			} else {
				mp[strRuneArr[rightIndex]] = rightIndex
			}
		}
		if rightIndex-leftIndex > ans {
			// 使用滑动窗口的左右下标定位数组
			ans = rightIndex - leftIndex
		}
	}
	return ans
}
