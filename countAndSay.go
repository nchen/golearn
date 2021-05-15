package main

import "strconv"

// https://leetcode-cn.com/problems/count-and-say/submissions/
func countAndSay(n int) string {
	var result string
	for i := 1; i <= n; i++ {
		if i == 1 {
			result = "1"
		} else {
			new_result := ""
			current_char := ""
			count := 0
			for j := 0; j < len(result); j++ {
				c := string(byte(result[j]))
				if c != current_char {
					if current_char != "" {
						new_result += strconv.Itoa(count) + current_char
					}
					current_char = c
					count = 1
				} else {
					count += 1
				}
			}
			if current_char != "" {
				new_result += strconv.Itoa(count) + current_char
			}
			result = new_result
		}
	}
	return result
}
