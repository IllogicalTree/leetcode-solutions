// 0001 - Two Sum
// https://leetcode.com/problems/two-sum/

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		if v, found := m[target-num]; found {
			return []int{v, i}
		}
		m[num] = i
	}
	return nil
}
