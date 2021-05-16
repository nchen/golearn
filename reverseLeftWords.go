// https://leetcode-cn.com/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof/submissions/
package main

func reverseLeftWords(s string, n int) string {
	return s[n:] + s[0:n]
}
