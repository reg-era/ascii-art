package ascii

import (
	"fmt"
	"os"
	"strings"
)

var Danger string = "❌ \033[31m"

func AsciiProcess(word string) {
	if len(word) == 0 {
		return
	}
	tabword := strings.Split(word, "\\n")
	for i := 0; i < len(tabword); i++ {
		if (len(tabword[i]) == 0) || (len(tabword[i]) == 0 && i != len(tabword)-1 && len(tabword[i+1]) == 0) {
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
			if word[j] < ' ' || word[j] > '~' {
				fmt.Println((Danger + "Character not found :("))
				os.Exit(0)
			}
			line := AsciiMap[word[j]][i]
			fmt.Print(line)
		}
		fmt.Println()
	}
}
