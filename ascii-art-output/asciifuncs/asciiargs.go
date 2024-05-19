package ascii

import (
	"fmt"
	"os"
)

func AsciiBannerSearch(arg string) string {
	filePrint := ""
	if arg == "standard" || arg == "thinkertoy" || arg == "shadow" || arg == "money" {
		arg += ".txt"
	}
	if arg == "standard.txt" || arg == "thinkertoy.txt" || arg == "shadow.txt" || arg == "money.txt" {
		filePrint = arg
	}
	return filePrint
}

func AsciiFilePrintSearch(arg string) string {
	file := ""
	if len(arg) > 9 && arg[:9] == "--output=" {
		file = arg[9:]
	}
	return file
}

func Error() {
	fmt.Println("Usage: go run . [STRING] [BANNER] [OPTION]\n\nExample: go run . something standard --output=<fileName.txt>")
	os.Exit(0)
}
