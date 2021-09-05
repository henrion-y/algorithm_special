package main

//给你一个字符串 s ，仅反转字符串中的所有元音字母，并返回结果字符串。
//
// 元音字母包括 'a'、'e'、'i'、'o'、'u'，且可能以大小写两种形式出现。
//
//
//
// 示例 1：
//
//
//输入：s = "hello"
//输出："holle"
//
//
// 示例 2：
//
//
//输入：s = "leetcode"
//输出："leotcede"
//
//
//
// 提示：
//
//
// 1 <= s.length <= 3 * 10⁵
// s 由 可打印的 ASCII 字符组成
//
// Related Topics 双指针 字符串 👍 220 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
// 这里数据不是很多， 不要用 map ，直接遍历性能更好！！
func isVowels(ch byte) bool {
	vowels := "aeiouAEIOU"
	for i := 0; i < len(vowels); i++ {
		if ch == vowels[i] {
			return true
		}
	}
	return false
}

// 还是双指针夹逼
func reverseVowels(s string) string {
	left, right := 0, len(s)-1
	s2 := []byte(s)
	for left < right {
		if isVowels(s2[left]) && isVowels(s2[right]) {
			s2[left], s2[right] = s2[right], s2[left]
			left++
			right--
		} else if !isVowels(s2[left]) && isVowels(s2[right]) {
			left++
		} else if isVowels(s2[left]) && !isVowels(s2[right]) {
			right--
		} else {
			left++
			right--
		}
	}
	return string(s2)
}

//leetcode submit region end(Prohibit modification and deletion)
