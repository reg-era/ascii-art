package ascii

import (
	"fmt"
	"os"
)

func AsciiCorrect(word string) {
	sliceword := []rune(word)
	for i := 0; i < len(sliceword); i++ {
		if sliceword[i] == 92 && sliceword[i+1] == 110 {
			sliceword[i] = '\n'
			sliceword = append(sliceword[:i+1], sliceword[i+2:]...)
			i--
		}
	}
	AsciiPrint(sliceword)
}

func AsciiPrint(tabword []rune) {
	if len(tabword) == 0 {
		os.Exit(0)
	}
	fmt.Println(tabword)
	for i := 0; i < 8; i++ {
		for j := 0; j < len(tabword); j++ {
			if tabword[j] == '\n' {
				fmt.Println()
				AsciiPrint(tabword[j+1:])
				break
			}
			line := AsciiMap[tabword[j]][i]
			fmt.Print(line)

		}
		//fmt.Println()
	}
}