package main

/**
数组中占比超过一半的元素称之为主要元素。给定一个整数数组，找到它的主要元素。若没有，返回-1。
https://leetcode-cn.com/problems/find-majority-element-lcci/
*/

func majorityElement(nums []int) int {
	size := len(nums)
	m := make(map[int]int)

	for _, v := range nums {
		m[v] += 1
	}

	for k, v := range m {
		if v > size/2 {
			return k
		}
	}
	return -1
}
