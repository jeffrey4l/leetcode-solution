package solution

// Question: https://leetcode.com/problems/palindrome-number/

import (
	"strconv"
)

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func isPalindrome2(x int) bool {
	if x < 0 || x%10 == 0 && x != 0 {
		return false
	}
	reverse := 0
	for x > reverse {
		reverse = reverse*10 + x%10
		x = x / 10
	}

	return x == reverse || x == reverse/10
}
