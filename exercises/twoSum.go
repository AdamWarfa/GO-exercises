package exercises

import "fmt"

func TwoSum(nums []int, target int) []int {
	result := []int{}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i]+nums[j] == target && i != j {
				result = []int{i, j}
			}
		}
	}
	fmt.Println(result)
	return result
}
