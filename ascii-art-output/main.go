package main

import (
	"flag"
	"fmt"
	"os"

	ascii "ascii/asciifuncs"
)

func main() {
	args := os.Args[1:]

	word := ""
	banner := ""
	filePrint := ""

	fileOption := flag.String("output", "def", "output usage")
	flag.Parse()

	switch len(args) {
	case 1:
		word = args[0]
		ascii.AsciiReload("standard.txt", word)
		ascii.AsciiCreat(ascii.AsciiProcess(word), filePrint)
	case 2:
		if *fileOption == "def" {
			word = args[0]
			banner = args[1]
		} else {
			filePrint = *fileOption
			word = args[1]
			banner = "standard.txt"
		}
		ascii.AsciiReload(ascii.AsciiSearch(banner), word)
		ascii.AsciiCreat(ascii.AsciiProcess(word), filePrint)
	case 3:
		word := args[1]
		banner := args[2]
		filePrint := *fileOption
		ascii.AsciiReload(ascii.AsciiSearch(banner), word)
		ascii.AsciiCreat(ascii.AsciiProcess(word), filePrint)
	default:
		fmt.Println("\033[31mUsage: go run . [STRING] [BANNER] [OPTION]\nExample: go run . something standard --output=<fileName.txt>")
		os.Exit(0)
	}
}
