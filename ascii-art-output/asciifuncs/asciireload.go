package ascii

import (
	"bufio"
	"fmt"
	"os"
)

var AsciiMap = map[rune][]string{}

func MapReload(scanner *bufio.Scanner) {
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
			AsciiMap[rune(count)] = append(AsciiMap[rune(count)], asciichar...)
			count++
		}
	}
}

func AsciiReload(filePrint string) {
	asciifile, err := os.Open("asciilib/" + filePrint)
	if err != nil {
		fmt.Println("Dont play in program file >(")
		os.Exit(0)
	}
	defer asciifile.Close()
	scanne := bufio.NewScanner(asciifile)
	MapReload(scanne)
}

func AsciiSearch(filePrint string) string {
	if filePrint != "standard.txt" && filePrint != "thinkertoy.txt" && filePrint != "shadow.txt" && filePrint != "standard" && filePrint != "thinkertoy" && filePrint != "shadow" && filePrint != "money" && filePrint != "money.txt" {
		fmt.Println(Danger + "ASCII-ART library not found")
		os.Exit(0)
	}
	if filePrint == "standard" || filePrint == "thinkertoy" || filePrint == "shadow" || filePrint == "money" {
		filePrint += ".txt"
	}
	return filePrint
}
