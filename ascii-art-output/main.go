package main

import (
	"os"

	ascii "ascii/asciifuncs"
)

func main() {
	args := os.Args[1:]

	word := ""
	banner := ""
	filePrint := ""

	// fileOption := flag.String("output", "def", "output usage")
	// flag.Parse()

	switch len(args) {
	case 1:
		word = args[0]
		ascii.AsciiReload("standard.txt", word)
		ascii.AsciiCreat(ascii.AsciiProcess(word), filePrint)
	case 2:
		word = args[0]
		banner = ascii.AsciiBannerSearch(args[1])
		if banner == "" {
			banner = "standard.txt"
			filePrint = ascii.AsciiFilePrintSearch(args[1])
			if filePrint == "" {
				ascii.Error()
			}
		}
		ascii.AsciiReload(banner, word)
		ascii.AsciiCreat(word, filePrint)
	case 3:
		word := args[1]
		banner = ascii.AsciiBannerSearch(args[1])
		filePrint = ascii.AsciiFilePrintSearch(args[2])
		if banner == "" {
			banner = ascii.AsciiBannerSearch(args[2])
			filePrint = ascii.AsciiFilePrintSearch(args[1])
			if filePrint == "" {
				ascii.Error()
			}
		}
		ascii.AsciiReload(banner, word)
		ascii.AsciiCreat(ascii.AsciiProcess(word), filePrint)
	default:
		ascii.Error()
	}
}
