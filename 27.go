package main

func removeElement(nums []int, val int) int {
    index := 0
	for i := 0; i < len(nums); i++ {
		if nums[index] != nums[i] {
			nums[index] = nums[i]
			index++
		}
	}
	return index
}