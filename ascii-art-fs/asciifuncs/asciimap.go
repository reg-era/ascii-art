package ascii

import (
	"bufio"
)

var AsciiMap = map[rune][]string{}

func MapReload(scanner *bufio.Scanner, word string) {
	count := 31
	for scanner.Scan() {
		asciichar := []string{}
		line := scanner.Text()
		if line == "" {
			count++
			continue
		} else {
			for i := 0; i < 8 && scanner.Scan(); i++ {
				asciichar = append(asciichar, line)
				line = scanner.Text()
				if i == 6 && count == 126 {
					asciichar = append(asciichar, "      ")
				}
			}
			if CheckChar(rune(count), word) {
				AsciiMap[rune(count)] = append(AsciiMap[rune(count)], asciichar...)
			}
			count++
		}
	}
}

func CheckChar(r rune, text string) bool {
	for _, char := range text {
		if char == r {
			return true
		}
	}
	return false
}
