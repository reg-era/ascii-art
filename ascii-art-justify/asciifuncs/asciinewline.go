package ascii

import "strings"

func CheckNewLine(word string) bool {
	check := len(word) / 2
	if len(word)%2 != 0 {
		return false
	}
	count := 0
	tab := strings.Split(word, "\\n")
	for i := 0; i < len(tab)-1; i++ {
		if tab[i] == "" {
			count++
		}
	}
	if count == check {
		return true
	} else {
		return false
	}
}
