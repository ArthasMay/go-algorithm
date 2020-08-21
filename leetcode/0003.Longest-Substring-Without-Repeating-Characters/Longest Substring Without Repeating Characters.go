package leetcode

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	var bitSet [256]uint8
	result, left, right := 0, 0, 0
	for left < len(s) {
		if right < len(s) && bitSet[s[right]] == 0 {
			bitSet[s[right]] = 1
			right++
		} else {
			bitSet[s[right]] = 0
			left++
		}
		result = max(result, right-left)
	}
	return result
}

func lengthOfLongestSubstring_(s string) int {
	if len(s) == 0 {
		return 0
	}

	var freq [256]int
	result, left, right := 0, 0, -1
	for left < len(s) {
		if right+1 < len(s) && freq[s[right+1]-'a'] == 0 {
			freq[s[right+1]]++
			right++
		} else {
			freq[s[right+1]]--
			left++
		}
		result = max(result, right-left+1)
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
