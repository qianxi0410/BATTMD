package hot100

import "sort"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	if len(nums1) & 1 == 1 {
		return float64(nums1[(len(nums1)-1)/2])
	}
	return float64(nums1[len(nums1)/2] + nums1[len(nums1)/2+1]) / 2
}
