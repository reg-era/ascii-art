package ascii

import (
	"fmt"
	"os"
	"strings"
)

var Danger string = "❌ \033[31m"

func AsciiCreat(cont, fileOption string) {
	if fileOption == "" {
		fmt.Print(cont)
		os.Exit(0)
	}
	if len(fileOption) > 4 {
		if fileOption[len(fileOption)-4:] != ".txt" {
			fmt.Println(Danger + "Something wrong check your files :)\nThe output file must be .txt extension")
			os.Exit(0)
		}
	} else {
		fmt.Println(Danger + "Something wrong check your files :)\nThe output file must be .txt extension")
		os.Exit(0)
	}
	os.WriteFile(fileOption, []byte(cont), 0o777)
}

func AsciiProcess(word string) string {
	res := ""
	if len(word) == 0 {
		return res
	}

	if len(word) != 1 && CheckNewLine(word) {
		for i := 0; i < len(word)/2; i++ {
			res += "\n"
		}
		return res
	} else {
		tabword := strings.Split(word, "\\n")
		for i := 0; i < len(tabword); i++ {
			if (len(tabword[i]) == 0) || (len(tabword[i]) == 0 && i != len(tabword)-1 && len(tabword[i+1]) == 0) {
				res += "\n"
				continue
			} else {
				tab := []rune(tabword[i])
				res += AsciiPrint(tab)
			}
		}
	}
	return res
}

func AsciiPrint(word []rune) string {
	res := ""
	if len(word) == 0 {
		return "\n"
	}
	for i := 0; i < 8; i++ {
		for j := 0; j < len(word); j++ {
			if word[j] < ' ' || word[j] > '~' {
				fmt.Println(Danger + "Character not found :(")
				os.Exit(0)
			}
			line := AsciiMap[word[j]][i]
			res += line
		}
		res = res + "\n"
	}
	return res
}
