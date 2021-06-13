// 0006 - ZigZag Conversion
// https://leetcode.com/problems/zigzag-conversion/

package leetcode

func convert(s string, numRows int) string {

	if numRows == 1 {
		return s
	}

	ans := ""
	ansMap := make([][]byte, numRows)
	index := 0
	goDown := true

	for i := 0; i < len(s); i++ {

		ansMap[index] = append(ansMap[index], s[i])

		if index == numRows-1 {
			goDown = false
		}
		if index == 0 {
			goDown = true
		}

		if goDown {
			index++
		} else {
			index--
		}
	}

	for i := 0; i < len(ansMap); i++ {
		ans += string(ansMap[i])
	}

	return ans
}
