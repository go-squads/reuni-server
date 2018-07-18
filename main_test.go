package main

import (
	"testing"
)

func TestMultiply(t *testing.T) {
	if(multiply(4,5) != 20) {
		t.Fail()
	}
}