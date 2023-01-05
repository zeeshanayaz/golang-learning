package main

import "log"

func main() {
	result := twoSum([]int{3, 2, 3}, 6)

	log.Println("Two Sums indexes are", result)
}

func twoSumSequential(nums []int, target int) []int {

	var result []int

	for key, value := range nums {
		if key != len(nums)-1 {
			if target == value+nums[key+1] {
				result = append(result, key)
				result = append(result, key+1)
				break
			}
		}
	}

	return result
}

func twoSum(nums []int, target int) []int {

	var result []int

	for key, value := range nums {
		if key != len(nums)-1 {
			for i := key + 1; i < len(nums); i++ {
				if target == value+nums[i] {
					result = append(result, key)
					result = append(result, i)
					break
				}
			}
		}
	}

	return result
}
