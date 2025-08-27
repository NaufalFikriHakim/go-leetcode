package main

import "sort"

func merge(nums1 []int, m int, nums2 []int, n int) {
    index := 0
    for index < n {
        nums1[m+index] = nums2[index]
        index++
    } 
	sort.Ints(nums1)
}