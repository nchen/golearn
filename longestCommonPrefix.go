// https://leetcode-cn.com/problems/longest-common-prefix/submissions/
package main

func longestCommonPrefix(strs []string) string {
	L := len(strs)
	if L == 0 {
		return ""
	}
	if L == 1 {
		return strs[0]
	}
	common := strs[0]
	for i := 1; i < L; i++ {
		s := strs[i]
		matchIndex := -1
		for j := 0; j < len(common); j++ {
			if j < len(s) && common[j] == s[j] {
				matchIndex = j
			} else {
				break
			}
		}
		common = common[:(matchIndex + 1)]
	}
	return common
}
