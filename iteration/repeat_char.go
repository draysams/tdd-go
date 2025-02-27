package iteration

import "strings"

const repeatCount = 5

func RepeatChar(character string) string {
	var repeatedString strings.Builder
	for range repeatCount {
		repeatedString.Write([]byte(character))
	}
	return repeatedString.String()
}
