package helpers

func RemoveQuotes(str string) string {
	strRunes := []rune(str)
	quoteRune := []rune("\"")[0]
	if strRunes[0] == quoteRune && strRunes[len(strRunes)-1] == quoteRune {
		return string(strRunes[1 : len(strRunes)-1])
	}

	return string(strRunes)
}
