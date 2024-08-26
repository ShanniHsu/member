package word

// 判斷一個字符串s是否時回文字符串
// 回文字符串是指正序(從左到右)倒序(從右到左)讀都是一樣的字符
func IsPalindrome(s string) bool {
	for i := range s {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
