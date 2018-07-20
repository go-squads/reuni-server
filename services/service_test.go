package services

import (
	"testing"
)

func TestTokenRandomizerDifferentAtLeastAHundredThousandTry(t *testing.T) {
	var data map[string]bool
	for i := 0; i < 100000; i++ {
		if data[generateToken()] {
			t.Fail()
		}
	}
}
