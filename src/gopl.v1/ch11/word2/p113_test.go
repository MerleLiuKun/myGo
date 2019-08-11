package word2

import (
	"math/rand"
	"testing"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randNonPalindrome(rng *rand.Rand) string {
	n := rng.Intn(23)
	n += 2
	myBytes := make([]byte, n)

	for i := 0; i < (n+1)/2; i++ {
		r := rng.Intn(len(charset) - 1)
		myBytes[i] = charset[r]
		myBytes[n-i-1] = charset[r+1]
	}
	return string(myBytes)
}

func TestRandomPalindrome(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed :%d", seed)

	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randNonPalindrome(rng)
		if isPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = True", p)
		}
	}
}
