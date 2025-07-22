package utils

func IsPalindrome(s string) bool {
	n := len(s)
	if n == 0 {
		return false
	}
	if n == 1 {
		return true
	}

	left := 0
	right := len(s) - 1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
