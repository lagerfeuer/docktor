package util

import "strings"

// RightPad2Len https://github.com/git-time-metric/gtm/blob/019de991bfe05a2a643d2a3bb9802ee9d3023a1d/util/string.go#L55-L69
func RightPad2Len(s string, padStr string, overallLen int) string {
	var padCountInt = 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = s + strings.Repeat(padStr, padCountInt)
	return retStr[:overallLen]
}

// LeftPad2Len https://github.com/git-time-metric/gtm/blob/019de991bfe05a2a643d2a3bb9802ee9d3023a1d/util/string.go#L55-L69
func LeftPad2Len(s string, padStr string, overallLen int) string {
	var padCountInt = 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = strings.Repeat(padStr, padCountInt) + s
	return retStr[(len(retStr) - overallLen):]
}
