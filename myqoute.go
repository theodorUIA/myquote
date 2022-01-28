package myqoute

import (
	"math/rand"

	"time"

	"rsc.io/quote"
)

func RandNum(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func GetQuote() string {
	min := 1
	max := 4
	var randNum = RandNum(min, max)
	newquote := "Failed from start."
	switch randNum {
	case 1:
		newquote = quote.Glass()
	case 2:
		newquote = quote.Go()
	case 3:
		newquote = quote.Opt()
	case 4:
		newquote = quote.Hello()
	default:
		newquote = "Failed to find number.."
	}
	return newquote

}
