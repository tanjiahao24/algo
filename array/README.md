# 数组
  

# 数组基础题目

## 保序数组
1. <font color=MediumSpringGreen>easy</font> [26. 删除有序数组中的重复项 removeDuplicates](https://leetcode.cn/problems/remove-duplicates-from-sorted-array/) 
2. <font color=MediumSpringGreen>easy</font> [283. 移动零 moveZeroes](https://leetcode.cn/problems/move-zeroes/description/)
3. <font color=MediumSpringGreen>easy</font> [88. 合并两个有序数组 merge](https://leetcode.cn/problems/merge-sorted-array/description/)

模版：
> 基本还是双指针，同向双指针，一前一后，根据前面的指针所对应的值；来判断后面的指针该如果运动
```go
count := 0
for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i] != nums[i-1] {
			nums[count] = nums[i]
			count++
		}
}
```