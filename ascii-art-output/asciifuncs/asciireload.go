package ascii

import (
	"bufio"
	"fmt"
	"os"
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

func AsciiReload(banner string, word string) {
	asciifile, err := os.Open("asciilib/" + banner)
	if err != nil {
		fmt.Println("Dont play in program file >(")
		os.Exit(0)
	}
	defer asciifile.Close()
	scanne := bufio.NewScanner(asciifile)
	MapReload(scanne, word)
}

func AsciiSearch(filePrint string) string {
	if filePrint == "standard" || filePrint == "thinkertoy" || filePrint == "shadow" || filePrint == "money" {
		filePrint += ".txt"
	}
	if filePrint != "standard.txt" && filePrint != "thinkertoy.txt" && filePrint != "shadow.txt" && filePrint != "money.txt" {
		fmt.Println(Danger + "ASCII-ART library not found")
		os.Exit(0)
	}
	return filePrint
}
