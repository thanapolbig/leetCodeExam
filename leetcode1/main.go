package main

import "log"

func main() {
	test := []int{2, 7, 11, 15}
	result := twoSum(test, 9)
	log.Println(result)
}
func twoSum(nums []int, target int) []int {
	for n, start := range nums {
		for i := 1; i < len(nums); i++ {
			if n == i {
				continue
			}
			if start+nums[i] == target {
				return []int{n, i}
			}
		}
	}
	return nil
}
