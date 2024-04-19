package client

func parseIntoCommandAndParams(rawCommand string) (string, []string) {
	wordList := []string{}
	newWord := ""
	isInQuotes := false

	for i := range len(rawCommand) {
		char := rawCommand[i : i+1]

		switch char {
		case "\"":
			if isInQuotes {
				isInQuotes = true
				wordList = append(wordList, newWord+char)
				newWord = ""
			} else {
				isInQuotes = false
				newWord += "\""
			}
		case " ":
			if isInQuotes {
				newWord += " "
			} else if newWord == "" {
				continue
			} else {
				wordList = append(wordList, newWord)
				newWord = ""
			}
		default:
			newWord += char
		}
	}

	if newWord != "" {
		wordList = append(wordList, newWord)
	}

	if len(wordList) == 0 {
		return "", []string{}
	}

	return wordList[0], wordList[1:]
}
