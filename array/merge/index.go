package merge

// 从末尾开始填空
func merge(nums1 []int, m int, nums2 []int, n int) {
	right1 := m - 1
	right2 := n - 1
	for len := m + n - 1; len >= 0; len-- {
		// 什么情况下需要填上 nums2 当前的值？
		// 当 right1 越界的时候，或者是 right1 没有越界 & right2 也没有越界 & right1 对应的值小于等于 right2 对应的值
		if right1 < 0 || (right2 >= 0 && nums1[right1] <= nums2[right2]) {
			nums1[len] = nums2[right2]
			right2--
		} else {
			nums1[len] = nums1[right1]
			right1--
		}
	}
}
