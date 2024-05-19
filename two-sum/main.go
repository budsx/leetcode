package main

import "fmt"

/*

Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.

Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

*/

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum2([]int{2, 7, 11, 15}, 9))
}

// 1. Create a map to store the value and index of the array
// 2. Loop through the array
// 3. Calculate the difference between the target and the current value
// 4. Check if the difference is in the map
// 5. If it is, return the index of the difference and the current index
// 6. If it is not, store the current value and index in the map
// 7. If no solution is found, return nil
// Iteration 1
// i = 0, diff = 9 - 2 = 7, maps[7] = 1
func twoSum(nums []int, target int) []int {
	maps := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		diff := target - nums[i]
		if c, ok := maps[diff]; ok {
			return []int{c, i}
		}
		maps[nums[i]] = i
	}
	return nil
}

// 1. Loop through the array
// 2. Loop through the array starting from the next index
// 3. Check if the sum of the two values is equal to the target
// 4. If it is, return the indices of the two values
// 5. If no solution is found, return nil
// Iteration 1 
// i = 0, j = 1, nums[0] + nums[1] = 2 + 7 = 9
func twoSum2(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
