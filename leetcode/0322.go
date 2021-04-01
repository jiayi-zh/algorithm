package leetcode

// f(n) = { f(n-coin) + 1      n >= 1
//        { 0                  n = 0
//        { -1                 n < 0 无解
//         n代表金额，也就是状态。 返回所需最少的硬币数
const INT_MAX = int(^uint(0) >> 1)

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1, amount+1)
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		minV := INT_MAX
		for _, coin := range coins {
			v := i - coin
			if v < 0 {
				continue
			} else if v == 0 {
				// 1 是最优解了
				minV = 1
				break
			} else {
				if dp[v] <= 0 {
					continue
				} else {
					if dp[v]+1 < minV {
						minV = dp[v] + 1
					}
				}
			}
		}
		if minV == INT_MAX {
			dp[i] = -1
		} else {
			dp[i] = minV
		}
	}
	return dp[amount]
}
