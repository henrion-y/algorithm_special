package main

import (
	"fmt"
	"strings"
)

//给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
//
// 说明：本题中，我们将空字符串定义为有效的回文串。
//
//
//
// 示例 1:
//
//
//输入: "A man, a plan, a canal: Panama"
//输出: true
//解释："amanaplanacanalpanama" 是回文串
//
//
// 示例 2:
//
//
//输入: "race a car"
//输出: false
//解释："raceacar" 不是回文串
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 2 * 10⁵
// 字符串 s 由 ASCII 字符组成
//
// Related Topics 双指针 字符串 👍 411 👎 0

//leetcode submit region begin(Prohibit modification and deletion)

func isalnum(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}

// 先遍历、创建新字符串， 再比较两个新字符串
func isPalindrome(s string) bool {
	var s2, s3 []byte
	for i := 0; i < len(s); i++ {
		if isalnum(s[i]) {
			s2 = append(s2, s[i])
		}
	}
	s2Str := strings.ToLower(string(s2))
	for i := len(s2Str) - 1; i >= 0; i-- {
		s3 = append(s3, s2Str[i])
	}
	return s2Str == string(s3)
}

// 双指针左右夹逼
func isPalindrome2(s string) bool {
	var s2 []byte
	for i := 0; i < len(s); i++ {
		if isalnum(s[i]) {
			s2 = append(s2, s[i])
		}
	}
	left, right := 0, len(s2)-1
	s2Str := strings.ToLower(string(s2))
	for left < right {
		if s2Str[left] != s2Str[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// 通过计算代替右指针
func isPalindrome3(s string) bool {
	var s2 []byte
	for i := 0; i < len(s); i++ {
		if isalnum(s[i]) {
			s2 = append(s2, s[i])
		}
	}
	s2Str := strings.ToLower(string(s2))
	n := len(s2Str)
	for i := 0; i < n/2; i++ {
		if s2Str[i] != s2Str[n-i-1] {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
// 测试用例:"A man, a plan, a canal: Panama" 测试结果:false 期望结果:true
func main() {
	fmt.Println(isPalindrome2("A man, a plan, a canal: Panama"))
}
