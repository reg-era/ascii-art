package ascii

import (
	"bufio"
)

var AsciiMap = map[rune][]string{}

func MapReload(scanner *bufio.Scanner) {
	//asciicharrange := []string{" /\\/| ", "|/\\/  ", "      ", "      ", "      ", "      ", "      ", "      "}
	count := 31
	for scanner.Scan() {
		asciichar := []string{}
		line := scanner.Text()
		if line == "" {
			count++
			continue
		} else {
			// if count == 126 {
			// 	AsciiMap[rune(count)] = append(AsciiMap[rune(count)], asciicharrange...)
			// 	break
			// }
			for i := 0; i < 8 && scanner.Scan(); i++ {
				asciichar = append(asciichar, line)
				line = scanner.Text()
			}
			AsciiMap[rune(count)] = append(AsciiMap[rune(count)], asciichar...)
			count++
		}
	}
}
