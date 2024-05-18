package main

import (
	"flag"
	"fmt"
	"os"

	ascii "ascii/asciifuncs"
)

var Danger string = "❌ \033[31m"

func main() {
	args := os.Args[1:]

	word := ""
	banner := "standard.txt"
	color := "white"
	spcolor := ""
	filePrint := ""

	optionfile := flag.String("output", "def", "output usage")
	optioncolor := flag.String("color", "def", "color usage")
	flag.Parse()
	switch len(args) {
	case 1:
		word = args[0]
		ascii.AsciiReload(banner, word)
		ascii.AsciiCreat(ascii.AsciiProcess(word, color, spcolor), filePrint)
	case 2:
		if ascii.IsBanner(args[1]) {
			banner = args[1]
			word = args[0]
		} else if *optioncolor != "def" && *optionfile != "def" {
			fmt.Println("sdsdscdscs")
			return
		} else {
			word = args[1]
			if *optioncolor != "def" {
				color = *optioncolor
			} else {
				filePrint = *optionfile
			}
		}
		ascii.AsciiReload(ascii.AsciiSearch(banner), word)
		ascii.AsciiCreat(ascii.AsciiProcess(word, color, spcolor), filePrint)
	case 3:
		if ascii.IsBanner(args[2]) && (*optioncolor == "def" || *optionfile == "def") {
			word = args[1]
			banner = args[2]
			if *optioncolor != "def" {
				color = *optioncolor
			} else {
				filePrint = *optionfile
			}
		} else if *optioncolor != "def" && *optionfile == "def" {
			word = args[2]
			spcolor = args[1]
			color = *optioncolor
		} else {
			word = args[2]
			color = *optioncolor
			filePrint = *optionfile
		}
		ascii.AsciiReload(ascii.AsciiSearch(banner), word)
		ascii.AsciiCreat(ascii.AsciiProcess(word, color, spcolor), filePrint)
	case 4:
		word := args[2]
		banner := args[3]
		color := *optioncolor
		filePrint := *optionfile
		ascii.AsciiReload(ascii.AsciiSearch(banner), word)
		ascii.AsciiCreat(ascii.AsciiProcess(word, color, spcolor), filePrint)
	
	default:
		fmt.Println("\033[31mUsage: go run . [STRING] [BANNER] [OPTION]\nExample: go run . something standard --output=<fileName.txt>")
		os.Exit(0)
	}
}
