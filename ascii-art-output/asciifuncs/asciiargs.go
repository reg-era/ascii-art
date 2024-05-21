package ascii

import (
	"fmt"
	"os"
	"strings"
)

// 3 variable that we need in this program
var Word, Banner, File string

// check each possible arg can be
func FiltringArgs(args []string) {
	switch len(args) {
	// in 1 arg must be only user word
	case 1:
		Word = args[0] // searching the word
	// in 2 arg must be user word and ( file output or banner name )
	case 2:
		for i := 0; i < 2; i++ {
			AsciiSearch(args[i]) // searching the args 2 time
		}
		if Word == "" && (File == "" || Banner == "") {
			Word = args[0]
			if Banner != "" {
				Banner = AsciiBannerSearch(args[1])
			}
		}
		if Word == "" || (File == "" && Banner == "") { // mean err detected
			Error("")
		}
	// in 3 must be word and output file and banner name
	case 3:
		for i := 0; i < 3; i++ {
			AsciiSearch(args[i]) // searching the args 3 time
		}
		// if we dont found word but we found option
		if Word == "" && (File != "" || Banner != "") {
			// in this case the word is obligate to be first argument
			if AsciiFilePrintSearch(args[0]) == "" { // we check if is not a output option its mean is a banner
				Word = args[0]
				if AsciiBannerSearch(args[1]) != "" {
					Banner = AsciiBannerSearch(args[1])
				}
				if AsciiFilePrintSearch(args[1]) != "" { // we check if is not a output option its mean is banner
					Banner = AsciiBannerSearch(args[2]) // we initialize the banner to the last argument
				}
			} else {
				Word = args[1]
				Banner = AsciiBannerSearch(args[2])
			}
		}
		if Word == "" || File == "" || Banner == "" { // mean err detected
			Error("")
		}
	default:
		Error("")
	}
}

// search for args value (banner name , user word , output file) and stocket in constents
func AsciiSearch(arg string) {
	if Word == "" {
		Word = AsciiWordSearch(arg)
	}
	if Banner == "" {
		Banner = AsciiBannerSearch(arg)
	}
	if File == "" {
		File = AsciiFilePrintSearch(arg)
	}
}

// receive arg and return the word if not return null
func AsciiWordSearch(arg string) string {
	res := ""
	// we make 2 test to check if the arg is not a banner name and is not a file output
	test1 := AsciiBannerSearch(arg)
	test2 := AsciiFilePrintSearch(arg)
	if test1 == "" && test2 == "" { // if test valid mean is a word so we return it
		res = arg
	}

	tabarg := strings.Split(arg, " ") // in case of many word in one arg we split it
	if len(tabarg) != 1 {             // if many word detected
		for i := 0; i < len(tabarg); i++ {
			ind := AsciiWordSearch(tabarg[i]) // we search if all word is not output option or banner name
			if ind == "" {
				Error("")
			}
		}
	}
	return res
}

// receive arg and return the banner name if not return null
func AsciiBannerSearch(arg string) string {
	filePrint := ""
	if arg == "standard" || arg == "thinkertoy" || arg == "shadow" || arg == "money" { // we correct the banner name if is not .txt
		arg += ".txt"
	}
	if arg == "standard.txt" || arg == "thinkertoy.txt" || arg == "shadow.txt" || arg == "money.txt" { // we check if arg is a banner if not return null
		filePrint = arg
	}
	return filePrint
}

// receive arg and return the file output if not return null
func AsciiFilePrintSearch(arg string) string {
	file := ""
	if len(arg) > 9 && arg[:9] == "--output=" { // we check if the first of this arg start with --output= wich mean is a flag if not return null
		file = arg[9:] // we return the rest of mean the name of output file
		// check file errors
		if arg == "--output=" || !strings.HasSuffix(arg, ".txt") || strings.Contains(arg, "/") || file == ".txt" {
			Error("check your file name\n")
		}
	}
	return file
}

// our error handler take specification as argument to add in the error program
func Error(specification string) {
	fmt.Println(specification + "Usage: go run . [STRING] [BANNER] [OPTION]\n\nExample: go run . something standard --output=<fileName.txt>")
	os.Exit(0)
}
