package leetcode

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	// k 是 Array的索引 index
	for k, v := range nums {
		if idx, ok := m[target-v]; ok {
			return []int{idx, k}
		}
		m[v] = k
	}
	return nil
}
