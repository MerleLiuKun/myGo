package word1

import "testing"

func TestPalindrome(t *testing.T) {
	// 单纯通过返回的结果进行判断
	if !isPalindrome("detartrated") {
		t.Error(`IsPalindrome("detartrated") = false`)
	}
	if !isPalindrome("kayak") {
		t.Error(`IsPalindrome("kayak") = false`)
	}
}

func TestNonPalindrome(t *testing.T) {
	if isPalindrome("palindrome") {
		t.Error(`IsPalindrome("palindrome") = true`)
	}
}

func TestFrenchPalindrome(t *testing.T) {
	if !isPalindrome("été") {
		t.Error(`IsPalindrome("été") = false`)
	}
}

func TestCanalPalindrome(t *testing.T) {
	input := "A man, a plan, a canal: Panama"
	if !isPalindrome(input) {
		t.Errorf(`IsPalindrome(%q) = false`, input)
	}
}
