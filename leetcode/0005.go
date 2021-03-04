package leetcode

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	charArr := []rune(s)
	dp := make([][]int, len(charArr))
	for i := range dp {
		dp[i] = make([]int, len(charArr))
	}
	for i := 0; i < len(charArr); i++ {
		dp[i][i] = 1
	}

	// 存在只有两个不相同元素的时候, 默认取第一个
	ansLeft, ansLen := 0, 1
	for j := 1; j < len(charArr); j++ {
		for i := 0; i < j; i++ {
			if charArr[i] == charArr[j] {
				// j-1-(i+1)+1 < 2 ==> j-i < 3 区间大小小于2，也就是说只有一个元素的时候也是成立的
				if j-i < 3 {
					dp[i][j] = 1
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}

			if dp[i][j] > 0 && j-i+1 > ansLen {
				ansLeft = i
				ansLen = j - i + 1
			}
		}
	}
	return string(charArr[ansLeft : ansLeft+ansLen])
}
