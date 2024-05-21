package main

import (
	"os"

	ascii "ascii/asciifuncs"
)

func main() {
	// Reading argument
	args := os.Args[1:]
	// Argument must be atleas 3
	if len(args) > 3 {
		ascii.Error("")
	}
	//collecting and filtring args
	ascii.FiltringArgs(args)

	// reload map with the nesecere banner
	ascii.AsciiReload(ascii.Banner, ascii.Word)
	// creat or print the user input
	ascii.AsciiCreat(ascii.AsciiProcess(ascii.Word), ascii.File)
}
