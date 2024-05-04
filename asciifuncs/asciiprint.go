package ascii

import (
	"fmt"
	"strings"
)

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

func AsciiCorrect(word string) {
	if len(word) == 0 {
		return
	}
	res := strings.Split(word, "\\n")
	AsciiProcess(res)
}

func AsciiProcess(tabword []string) {
	for i := 0; i < len(tabword); i++ {
		if len(tabword[i]) == 0 && i != len(tabword)-1 && len(tabword[i+1]) == 0 {
			fmt.Println()
			continue
		}
		if len(tabword[i]) == 0 {
			fmt.Println()
			continue
		} else {
			tab := []rune(tabword[i])
			AsciiPrint(tab)
		}
	}
}

func AsciiPrint(word []rune) {
	if len(word) == 0 {
		fmt.Println()
		return
	}
	for i := 0; i < 8; i++ {
		for j := 0; j < len(word); j++ {
			line := AsciiMap[word[j]][i]
			fmt.Print(line)
		}
		fmt.Println()
	}
}
