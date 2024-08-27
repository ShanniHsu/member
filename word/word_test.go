package word

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//參考https://juejin.cn/post/6908938380114034701

// 使用第三方斷言庫 testify
func TestIsPalindrome(t *testing.T) {
	// 斷言IsPalindrome方法的返回值為True
	assert.True(t, IsPalindrome("detartrated"))
	assert.True(t, IsPalindrome("kayak"))
}

func TestNonPalindrome(t *testing.T) {
	// 斷言IsPalindrome方法的返回值為False
	assert.False(t, IsPalindrome("palindrome"))
}

//用go的test
//func TestIsPalindrome(t *testing.T) {
//	//因為是回文字符串，如果是false代表有問題
//	if !IsPalindrome("detartrated") {
//		t.Error(`IsPalindrome("detartrated") = false`)
//	}
//
//	if !IsPalindrome("kayak") {
//		t.Error(`IsPalindrome("kayak") = false`)
//	}
//}
//
//func TestNonPalindrome(t *testing.T) {
//	//這不是回文字符串，如果是true代表有問題
//	if IsPalindrome("palindrome") {
//		t.Error(`IsPalindrome("palindrome") = true`)
//	}
//}
