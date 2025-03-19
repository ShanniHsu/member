package exame

// 判斷回文字符串
func IsPalindrome(s string) bool {
	runes := []rune(s)
	for i := 0; i < len(runes)/2; i++ {
		if runes[i] != runes[len(runes)-1-i] {
			return false
		}
	}
	return true
}

// 兩數之和
func TwoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// 反轉字符串
func ReverseString(s string) string {
	runes := []rune(s)
	var reverse []rune
	for i := 0; i < len(runes); i++ {
		x := runes[len(runes)-1-i]
		reverse = append(reverse, x)
	}
	return string(reverse)
}

// 計算質數
// 檢查一個數字是否為質數。質數是指只能被 1 和自身整除的數字。
func IsPrime(n int) bool {
	switch {
	case n <= 1:
		return false

	case n == 2:
		return true

	case n%2 == 0:
		return false

	case n > 2:
		for i := n; i < n; i++ {
			if n%i == 0 {
				return false
			}
		}
		return true
	}
	return false
}

// 計算斐波那契數列
// 參考來源: http://www.mathsgreat.com/fibon/fibon_123.pdf
func Fib(n int) int {
	// 實作程式碼

}
