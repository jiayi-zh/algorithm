package leetcode

// 利用 哈希 来保证结果唯一
func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)
	for i, v := range nums {
		if r, ok := mp[target-v]; ok {
			return []int{i, r}
		} else {
			mp[v] = i
		}
	}
	return nil
}
