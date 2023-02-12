package helpers

func RemoveQuotes(str string) string {
	strRunes := []rune(str)
	quotesRunes := []rune("\"")
	if strRunes[0] == quotesRunes[0] && strRunes[len(strRunes)-1] == quotesRunes[0] {
		return string(strRunes[1 : len(strRunes)-1])
	}

	return string(strRunes)
}
