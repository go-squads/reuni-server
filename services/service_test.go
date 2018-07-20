package services

import (
	"testing"
)

func TestTokenRandomizerDifferentAtLeastAHundredThousandTry(t *testing.T) {
	var data map[string]bool
	data = make(map[string]bool)
	for i := 0; i < 100000; i++ {
		token := generateToken()
		if data[token] {
			t.Fail()
		} else {
			data[token] = true
		}
	}
}
