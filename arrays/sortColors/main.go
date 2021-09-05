package main

import (
	"fmt"
)

//给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
//
// 此题中，我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。
//
//
//
//
//
//
// 示例 1：
//
//
//输入：nums = [2,0,2,1,1,0]
//输出：[0,0,1,1,2,2]
//
//
// 示例 2：
//
//
//输入：nums = [2,0,1]
//输出：[0,1,2]
//
//
// 示例 3：
//
//
//输入：nums = [0]
//输出：[0]
//
//
// 示例 4：
//
//
//输入：nums = [1]
//输出：[1]
//
//
//
//
// 提示：
//
//
// n == nums.length
// 1 <= n <= 300
// nums[i] 为 0、1 或 2
//
//
//
//
// 进阶：
//
//
// 你可以不使用代码库中的排序函数来解决这道题吗？
// 你能想出一个仅使用常数空间的一趟扫描算法吗？
//
// Related Topics 数组 双指针 排序 👍 993 👎 0

/*
先把两个问题分解成一个小问题， 然后先解决这个小问题， 再用同样的方法解决另一个小问题
*/
func swapColors(colors []int, target int) (countTarget int) {

	for i, c := range colors {
		if c == target {
			colors[i], colors[countTarget] = colors[countTarget], colors[i]
			countTarget++
		}
	}
	return

}

func sortColors(nums []int) {
	countTarget := swapColors(nums, 0)
	swapColors(nums[countTarget:], 1)
}

// 测试用例:[2,0,2,1,1,0] 测试结果:[0,0,1,2,2,1] 期望结果:[0,0,1,1,2,2]
/*
2 0 2 1 1 0
0 2 2 1 1 0
	i=1 p0=1 p1-1
0 1 2 2 1 0
	i=3 p0=1 p1=2
0 1 1 2 2 0
	i=4 p0=1 p1=3
0 0 1 2 2 1
	i=5 p0=2 p1=4
0 0 1 1 2 2
	i=5 p0=

 */
// 荷兰国旗
func sortColors2(nums []int) {
	p0, p1 := 0, 0
	for i, c := range nums {
		if c == 0 {
			// todo
			nums[i], nums[p0] = nums[p0], nums[i]			// 交换完成后p0位置是0，i位置可能是1也可能是2
			if p0 < p1 {				// 此时p1走在p0前面，p1后面的都是1， 而i又在p1前面
				nums[i], nums[p1] = nums[p1], nums[i]
			}
			// 这里p0和p1都要走两步可以理解，但是p1为什么要跟i交换呢？
			p0++
			p1++
		} else if c == 1 {
			nums[i], nums[p1] = nums[p1], nums[i]
			p1++
		}
	}
}

// 两边向中间靠拢
/*
将问题分成两步
1、从右边找一个不是2的元素跟左边的元素对调，缩小右边小数组范围
2、在右边小数组中将0移动到前面。
 */
// todo
func sortColors3(nums []int) {
	p0, p2 := 0, len(nums)-1
	for i := 0; i <= p2; i++ {

		// 从右边找一个不是2的元素跟左边的元素对调，缩小右边小数组范围，
		// 也能保证当前小数组需要被调换的数字不是2
		for ; i <= p2 && nums[i] == 2; p2-- {
			nums[i], nums[p2] = nums[p2], nums[i]
		}
		// 在右边小数组中将0移动到前面
		if nums[i] == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			p0++
		}
	}
}


func main() {
	nums := []int{2,0,2,1,1,0}
	sortColors2(nums)
	fmt.Println(nums)

}
