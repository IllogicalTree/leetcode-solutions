// 0009 - Palindrome Number
// https://leetcode.com/problems/palindrome-number/

package leetcode

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	ans := 0
	tmp := x
	for x != 0 {
		ans = ans*10 + x%10
		x /= 10
	}
	return tmp == ans
}
