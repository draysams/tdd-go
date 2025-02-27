package iteration

import "strings"

func RepeatChar(character string, repeatCount int) string {
	var repeatedString strings.Builder
	for range repeatCount {
		repeatedString.Write([]byte(character))
	}
	return repeatedString.String()
}
