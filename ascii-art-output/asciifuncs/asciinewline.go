package ascii

import "strings"

// in case of only new lines in put
func CheckNewLine(word string) bool {
	// first we check if input contain only \n by counting \n
	check := len(word) / 2 // defind \n need
	count := 0             // newline counter
	tab := strings.Split(word, "\\n")
	for i := 0; i < len(tab)-1; i++ {
		if tab[i] == "" {
			count++
		}
	}
	if count == check { // if \n need same counter mean the input is only \n
		return true
	} else {
		return false
	}
}
