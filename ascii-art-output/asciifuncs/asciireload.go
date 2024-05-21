package ascii

import (
	"bufio"
	"os"
)

// the that we tack as ressource in this project take the character as keys and the slice of string(line) as content
var AsciiMap = map[rune][]string{}

// receive banner and the word to reload only char we need from the word
func AsciiReload(banner string, word string) {
	if banner == "" { // if ther is no banner it mean we work on standard
		banner = "standard.txt"
	}
	asciifile, err := os.Open("asciilib/" + banner) // we open the banner file
	if err != nil {
		Error("Dont play in program file >(\n")
	}
	defer asciifile.Close()               // closing file in last
	scanne := bufio.NewScanner(asciifile) // we tack content of file and put it in scanner
	MapReload(scanne, word)               // reloading asciimap
}

// take content of file and the user word
func MapReload(scanner *bufio.Scanner, word string) {
	count := 31          // we start with char behind space
	for scanner.Scan() { // if there is line to scan
		asciichar := []string{} // initialiw the content of keys map
		line := scanner.Text()
		if line == "" { // if empty line mean lets scan next rune else we are current scaning a rune
			count++
		} else {
			for i := 0; i < 8 && scanner.Scan(); i++ { // while there is 8 line and line to scan keep scan
				asciichar = append(asciichar, line)
				line = scanner.Text()
				if i == 6 && count == 126 { // check the last character and last line to append it
					asciichar = append(asciichar, line)
				}
			}
			if CheckChar(rune(count), word) { // if this rune is usfull in word ok append it in map
				AsciiMap[rune(count)] = append(AsciiMap[rune(count)], asciichar...)
			}
			count++ // increment the next rune
		}
	}
}

// check if the rune is in the user word
func CheckChar(r rune, word string) bool {
	for _, char := range word {
		if char == r {
			return true
		}
	}
	return false
}
