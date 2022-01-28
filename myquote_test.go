package myquote

import (
	"testing"
)

// Tester om RandNum() faktiske tall er mellom min og max
func TestRandNum(t *testing.T) {
	min := 1
	max := 4
	var randNum = RandNum(min, max)
	if randNum < min || randNum > max {
		t.Errorf("RandNum(%d, %d) = %d, want %d <= x <= %d", min, max, randNum, min, max)
	}
}

//getQuote() skal returnere en quote fra quote.go
func TestGetQuote(t *testing.T) {
	newquote := GetQuote()
	if newquote == "" {
		t.Errorf("GetQuote() = %s, want string", newquote)
	}
}
