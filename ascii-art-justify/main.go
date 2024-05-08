package main

import (
	"fmt"
	"os"

	ascii "ascii/asciifuncs"
)

var Danger string = "❌ \033[31m"

func main() {
	args := os.Args[1:]

	align := ""

	switch len(args) {
	case 0:
		fmt.Println(Danger + "Please insert text and the forme you want to print :)\nUsage: go run . [STRING] [BANNER] [OPTION]\nExample: go run . something standard --output=<fileName.txt>")
		os.Exit(0)
	case 1:
		fmt.Println(Danger + "Please inser the forme you want to print\nUsage: go run . [STRING] [BANNER] [OPTION]\nExample: go run . something standard --output=<fileName.txt>")
		os.Exit(0)
	case 2:
		word := args[0]
		filePrint := args[1]
		ascii.AsciiReload(ascii.AsciiSearch(filePrint))
		if len(word) != 1 && ascii.CheckNewLine(word) {
			for i := 0; i < len(word)/2; i++ {
				fmt.Println()
			}
			return
		} else {
			fmt.Print(ascii.AsciiProcess(word, "", ""))
		}
	case 3:
		word := args[0]
		filePrint := args[1]
		option := args[2]
		ascii.AsciiReload(ascii.AsciiSearch(filePrint))
		if len(option) >= 9 && (option[:8] == "--align=" || option[:7] == "-align=") {
			if option[:8] == "--align=" {
				align = option[8:]
			} else {
				align = option[7:]
			}

			ascii.AsciiReload(ascii.AsciiSearch(filePrint))
			if len(word) != 1 && ascii.CheckNewLine(word) {
				for i := 0; i < len(word)/2; i++ {
					fmt.Println()
				}
				return
			} else {
				fmt.Print(ascii.JustifyContent(align, word))
			}
		} else {
			fmt.Println(Danger + "Please inser the option you want to print\nUsage: go run . [STRING] [BANNER] [OPTION]\nExample: go run . something standard --output=<fileName.txt>")
			os.Exit(0)
		}
	default:
		fmt.Println(Danger + "Too many argument :)\nUsage: go run . [STRING] [BANNER] [OPTION]")
		os.Exit(0)
	}
}
