package movezeroes

func moveZeroes(nums []int) {
	len := len(nums)
	// 左右指针
	left := 0
	right := 0
	for right < len {
		if nums[right] != 0 { // 如果右指针的值为非零的时候有两种情况
			if left != right { // 1. 左右指针不在同一个位置的时候，将 left 对应的值置为 right 的值，right 置为 0
				nums[left] = nums[right]
				nums[right] = 0
			} // 2. 左右指针在同一个位置，直接都前进
			left++
			right++
		} else { // 如果右指针所对应的值为 0 的时候，右指针直接前进
			right++
		}
	}
}

// 方法2 遍历两遍 一边把非 0 的值放到前面来，第二遍把后面的值都置为 0
// func moveZeroes(nums []int) {
// 	len := len(nums)
//     n := 0
//     for i := 0; i < len; i++ {
//         if nums[i] != 0 {
//             nums[n] = nums[i]
//             n++
//         }
//     }
//     for n < len {
//         nums[n] = 0
//         n++
//     }
// }
