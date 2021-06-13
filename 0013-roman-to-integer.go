// 0013 - Roman to Integer
// https://leetcode.com/problems/roman-to-integer/

package leetcode

func romanToInt(s string) int {
	numerals := map[string]int{
		"I":  1,
		"IV": 4,
		"V":  5,
		"IX": 9,
		"X":  10,
		"XL": 40,
		"L":  50,
		"XC": 90,
		"C":  100,
		"CD": 400,
		"D":  500,
		"CM": 900,
		"M":  1000,
	}

	ans := 0
	i := 0

	for i < len(s) {
		// Handle 2 letter numerals
		if i < len(s)-1 && numerals[s[i:i+2]] != 0 {
			ans += numerals[s[i:i+2]]
			i += 2
		} else {
			ans += numerals[s[i:i+1]]
			i++
		}
	}
	return ans
}
