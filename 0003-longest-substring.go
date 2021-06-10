// 0003 - Longest Substring Without Repeating Characters
// https://leetcode.com/problems/longest-substring-without-repeating-characters/

package leetcode

func lengthOfLongestSubstring(s string) int {

	var longestLength int
	var index [128]int

	for left, right := 0, 0; right < len(s); right++ {
		if index[s[right]] > left {
			left = index[s[right]]
		}
		if right-left+1 > longestLength {
			longestLength = right - left + 1
		}
		index[s[right]] = right + 1
	}

	return longestLength
}
