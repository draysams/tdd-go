package iteration

const repeatCount = 5

func repeatChar(character string) string {
	repeatedString := ""
	for range repeatCount {
		repeatedString += character
	}
	return repeatedString
}
