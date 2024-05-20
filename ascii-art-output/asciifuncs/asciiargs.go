package ascii

import (
	"fmt"
	"os"
)

var Word, Banner, FilePrint string

func AsciiArgs(args []string) {
	if len(args) != 3 {
		Error()
	}

	Word = args[0]
	banner, berr := AsciiBannerSearch(args[1])
	filePrint, ferr := AsciiFilePrintSearch(args[2])
	if !berr || !ferr {
		Error()
	}
	Banner = banner
	FilePrint = filePrint
}

func AsciiBannerSearch(arg string) (string, bool) {
	if arg == "standard" || arg == "thinkertoy" || arg == "shadow" || arg == "money" {
		arg += ".txt"
	}

	if arg == "standard.txt" || arg == "thinkertoy.txt" || arg == "shadow.txt" || arg == "money.txt" {
		return arg, true
	} else {
		return arg, false
	}
}

func AsciiFilePrintSearch(arg string) (string, bool) {
	file := ""
	if len(arg) > 9 && arg[:9] == "--output=" {
		file = arg[9:]
	}

	if len(file) == 0 {
		return file, false
	} else {
		return file, true
	}
}

func Error() {
	fmt.Println("Usage: go run . [STRING] [BANNER] [OPTION]\n\nExample: go run . something standard --output=<fileName.txt>")
	os.Exit(0)
}
