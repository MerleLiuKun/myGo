package word2

import (
	"math/rand"
	"testing"
	"time"
)

func TestIsPalindrome(t *testing.T) {
	// 存储测试用例  结构体数组
	var tests = []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{input: "de", want: false},
		{"aa", true},
		{"kayak", true},
		{"detartrated", true},
		{"a man, a plan, a canal: Panama", true},
		{"Evil I did dwell; lewd did i live.", true},
		{"été", true},
		{"desserts", false},
	}

	for _, test := range tests {
		if got := isPalindrome(test.input); got != test.want {
			t.Errorf("IsPalindrome(%q) = %v", test.input, got)
		}
	}
}

func randomPalindrome(rng *rand.Rand) string {
	// 制定随机最大长度
	n := rng.Intn(25)
	runes := make([]rune, n)

	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000)) // 字符范围
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}

func TestIsPalindromeWithRandom(t *testing.T) {
	seed := time.Now().UTC().UnixNano() // 指定随机数种子
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !isPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}
