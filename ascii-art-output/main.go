package main

import (
	"os"

	ascii "ascii/asciifuncs"
)

func main() {
	args := os.Args[1:]

	ascii.AsciiArgs(args)

	ascii.AsciiReload(ascii.Banner, ascii.Word)
	ascii.AsciiCreat(ascii.AsciiProcess(ascii.Word), ascii.FilePrint)
}
