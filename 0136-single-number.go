// 0136 Single Number
// https://leetcode.com/problems/single-number/

package leetcode

func singleNumber(nums []int) int {
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		ans = ans ^ nums[i]
	}
	return ans
}
