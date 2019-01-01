// Package twosum problem: https://leetcode.com/problems/two-sum/solution/
package solution

func twoSum1(nums []int, target int) []int {

	for index, v := range nums {
		need := target - v
		for startIndex := index + 1; startIndex < len(nums); startIndex++ {
			if need == nums[startIndex] {
				return []int{index, startIndex}
			}
		}

	}
	return []int{}
}
func twoSum2(nums []int, target int) []int {

	indexes := make(map[int]int, len(nums))
	for index, v := range nums {
		indexes[v] = index
	}

	for index, v := range nums {
		needed := target - v
		if foundedIndex, ok := indexes[needed]; ok && foundedIndex != index {
			return []int{index, foundedIndex}
		}

	}
	return []int{}
}
func twoSum3(nums []int, target int) []int {

	indexes := make(map[int]int, len(nums))

	for index, v := range nums {
		if foundedIndex, ok := indexes[target-v]; ok {
			return []int{index, foundedIndex}
		}
		indexes[v] = index

	}
	return []int{}
}
