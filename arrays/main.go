package main

import (
	"fmt"
	"sort"
)

/*
数组和字符串基础的东西：
1、排序		--  快速排序

算法精要 ：
第一步、选定一个数， 将大于他的数移动到他右边， 小于他的数移动到他左边
第二步、使用分治对他左边、右边的数组重复以上操作
func quickSort(nums []int, left, right int) {
	// 1、选择一个基数， 为了方便移动， 选择中间位置的元素
	key := nums[(right+left)/2+1]
	// 2、复制左右游标 (这个游标在运算的时候需要移动，在分治的时候又要记录初始位置)
	i, j := left, right
	// 3、在左右两个游标会和的时候， 将大于key的数移动到key右边， 小于key的数移动到key左边
	for i <= j {
		for nums[i] < key {
			i++
		}
		for nums[j] > key {
			j--
		}
		if i <= j {			// 在左边找到一个大于基数的数、右边找到一个小于基数的数之后， 对调这两个数，然后移动左右游标
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
	}

	fmt.Println(key, nums)
	// 4、使用分治的将数组分成左右两部分，分别对左右两部分再做排序
	if left<j{
		quickSort(nums, left, j)
	}
	if right>i{
		quickSort(nums, i, right)
	}
}

2、查找		--	二分查找
算法精要：
需要通过修改左右两个游标来不断的缩小查找范围
然后每次获取两个游标中间的数字进行比较
func binarySearch(nums []int, v int) int {
	// 获取最左边和最右边两个下标
	left, right := 0, len(nums)-1
	// 当左右两个下标会和的时候， 说明所有情况都遍历完了
	for left <= right {
		// 获取中间下标的元素
		t := (right + left) / 2
		if nums[t] == v { // 如果中间位置元素等于V， 说明找到了，直接返回
			return t
		} else if nums[t] < v {
			left = t + 1 // 这里将左游标改为 t+1 。 注意， 这里t已经验证过了不是我们要的数据， 所以从t+1开始
		} else {
			right = t - 1 // 这里将右游标改为 t-1 。 注意， 这里t已经验证过了不是我们要的数据， 所以从t-1开始
		}
	}
	return -1
}

3、常见题目
a、删除有序数组重复项，返回删除后数组长度
注意审题， 【有序数组】
我们可以在通过下标遍历数组的时候对比是否跟上一个位置相等得出是否重复。
然后我们只要把非重复的保留到数组的前面就行了， 这样我们可以多定义一个游标short。
每次有不同的元素进来， 就放到short对应位置，再移动short
func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	short := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[short-1] {
			nums[short] = nums[i]
			short++
		}
	}
	return short
}

b、求两个有序数组的交集
数组是有序的，我们可以通过定义两个下标，通过下标的移动来判断元素是否相等
func intersect(nums1 []int, nums2 []int) []int {
	length1, length2 := len(nums1), len(nums2)
	index1, index2 := 0, 0
	var intersection []int
	for index1 < length1 && index2 < length2 { // 遍历较短的数组
		if nums1[index1] < nums2[index2] { // 如果短数组中当前游标所指向元素小于长数组中元素， 短数组游标后移以为
			index1++
		} else if nums1[index1] > nums2[index2] { // 如果短数组中当前游标所指向元素大于长数组中元素， 长数组游标后移以为
			index2++
		} else { // 短数组中游标所指元素等于长数组中所指， 则是重复的元素， 记录下来， 且两个游标都后移
			intersection = append(intersection, nums1[index1])
			index1++
			index2++
		}
	}
	return intersection
}

c、求两数之和
使用哈希表算法精要
1、只遍历一遍数组， 一边遍历一边往哈希表里面存数据
2、通过减法确定与当前元素匹配的另外一个元素的值， 再去哈希表中直接看改值是否存在
如果存在则返回， 不存在则将元素加入哈希表【省去了哈希表的去重】！！

func twoSum(nums []int, target int) []int {
	numSet := make(map[int]int)
	for i, v := range nums {
		if index, ok := numSet[target-v]; ok {
			return []int{index, i}
		}
		numSet[v] = i
	}
	return []int{}
}

d、求三数之和
算法精要：
先对数组进行排序，再遍历数组， 选择一个数作为基数，
然后通过左右两个下标遍历他后面的子数组， 采用左右夹逼的办法
关键在怎么对结果集进行去重， 这个还没搞清楚

func threeSum(nums []int) [][]int {
	if nums == nil || len(nums) < 3 {
		return [][]int{}
	}
	sort.Ints(nums)
	var result [][]int
	for k, v := range nums {
		if k > 0 && nums[k] == nums[k-1] { // 当遇到遍历的时候第一个数据数据重复了，那么就直接跳过
			continue
		}
		left, r := k+1, len(nums)-1
		for left < r {
			sum := nums[left] + nums[r] + v

			if sum == 0 {
				sum := []int{nums[left], nums[r], v}
				result = append(result, sum)

				for left < r && nums[left] == nums[left+1] { // 这里是观察后面的 left和r 往前和往后是否还是有重复的数据如果有的话直接就跳过即可
					left++
				}
				for left < r && nums[r] == nums[r-1] {
					r--
				}
				left++
				r--
			} else if sum < 0 {
				left++
			} else if sum > 0 {
				r--
			}
		}
	}
	return result
}
*/

// 求三数之和， 关键在怎么去重， 这个还没搞清楚
func threeSum(nums []int) [][]int {
	if nums == nil || len(nums) < 3 {
		return [][]int{}
	}
	var rest [][]int
	n := len(nums)
	sort.Ints(nums)
	// 这里存进去的时候还要对集合进行去重
	for index, v := range nums {
		// 这里解决了重复的问题
		if index > 0 && nums[index] == nums[index-1] {
			continue
		}
		left, right := index+1, n-1
		for left < right {
			sum := nums[left] + nums[right] + v
			if sum == 0 {
				rest = append(rest, []int{v, nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum > 0 {
				right--
			} else {
				left++
			}
		}

	}
	return rest
}

/*
算法精要：
先对数组进行排序，再遍历数组， 选择一个数作为基数，
然后通过左右两个下标遍历他后面的子数组， 采用左右夹逼的办法
关键在怎么对结果集进行去重， 这个还没搞清楚
*/
func threeSum2(nums []int) [][]int {
	if nums == nil || len(nums) < 3 {
		return [][]int{}
	}
	sort.Ints(nums)
	var result [][]int
	for k, v := range nums {
		if k > 0 && nums[k] == nums[k-1] { // 当遇到遍历的时候第一个数据数据重复了，那么就直接跳过
			continue
		}
		left, r := k+1, len(nums)-1
		for left < r {
			sum := nums[left] + nums[r] + v

			if sum == 0 {
				sum := []int{nums[left], nums[r], v}
				result = append(result, sum)

				for left < r && nums[left] == nums[left+1] { // 这里是观察后面的 left和r 往前和往后是否还是有重复的数据如果有的话直接就跳过即可
					left++
				}
				for left < r && nums[r] == nums[r-1] {
					r--
				}
				left++
				r--
			} else if sum < 0 {
				left++
			} else if sum > 0 {
				r--
			}
		}
	}
	return result
}

func main() {
	// 测试用例:[0,3,0,1,1,-1,-5,-5,3,-3,-3,0] 测试结果:[[0,3,-3],[0,0,0]] 期望结果:[[-3,0,3],[-1,0,1],[0,0,0]]
	nums := []int{0,0, 0, 0}
	fmt.Println(nums[:2])
	fmt.Println(threeSum(nums))

	fmt.Println(nil==[][]int{})
}
