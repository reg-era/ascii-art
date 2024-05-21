package ascii

import (
	"fmt"
	"os"
	"strings"
)

// printing each char word from map AsciiMap
func AsciiPrint(word []rune) string {
	res := ""
	if len(word) == 0 { // it mean \n
		return "\n"
	}
	for i := 0; i < 8; i++ { // point on each line
		for j := 0; j < len(word); j++ { // point on each char from the word
			if word[j] < ' ' || word[j] > '~' {
				Error("the character " + string(word[j]) + " can't be converted to ASCII art\n")
			}
			line := AsciiMap[word[j]][i] // print for each char the line that we point it
			res += line
		}
		res = res + "\n" // when we finish word we make new line
	}
	return res
}

// proccessing the word by checking \n and print it from map
func AsciiProcess(word string) string {
	res := ""
	// checking new lines input
	if len(word) != 1 && CheckNewLine(word) {
		for i := 0; i < len(word)/2; i++ { // print this input newlines
			res += "\n"
		}
		return res
	} else {
		tabword := strings.Split(word, "\\n") // split the word by \n and range on it
		for i := 0; i < len(tabword); i++ {
			if len(tabword[i]) == 0 { // check empty case mean newline
				res += "\n"
				continue
			} else { // else slice the case word and send it to print
				tab := []rune(tabword[i])
				res += AsciiPrint(tab)
			}
		}
	}
	return res
}

// creating file or printion result if filename null
func AsciiCreat(cont, file string) {
	if file == "" { // if file name is null mean we need to print it without creaing file
		fmt.Print(cont)
		os.Exit(0)
	}
	if len(file) > 4 { // check if extention file is .txt
		if file[len(file)-4:] != ".txt" {
			Error("The output file must be .txt extension\n")
		}
	}
	// creat file and put the content in it
	os.WriteFile(file, []byte(cont), 0o777)
}
