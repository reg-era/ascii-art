package main

import (
	"bufio"
	"fmt"
	"os"

	ascii "ascii/asciifuncs"
)

var Danger string = "❌ \033[31m"

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println(Danger + "Please insert text and the forme you want to print :)\nUsage: go run . [STRING] [BANNER]")
		os.Exit(0)
	} else if len(args) > 2 {
		fmt.Println(Danger + "Too many argument :)\nUsage: go run . [STRING] [BANNER]")
		os.Exit(0)
	} else if len(args) == 1 || len(args) == 2 {
		word := args[0]
		file := ""
		if len(args) == 2 {
			file = args[1]
		} else {
			file = "standard.txt"
		}

		if file != "standard.txt" && file != "thinkertoy.txt" && file != "shadow.txt" && file != "standard" && file != "thinkertoy" && file != "shadow" && file != "money" && file != "money.txt" {
			fmt.Println(Danger + "ASCII-ART library not found")
			os.Exit(0)
		}
		if file == "standard" || file == "thinkertoy" || file == "shadow" || file == "money" {
			file += ".txt"
		}
		asciifile, err := os.Open("asciilib/" + file)
		if err != nil {
			fmt.Println("Dont play in program file >(")
			os.Exit(0)
		}
		defer asciifile.Close()
		scanne := bufio.NewScanner(asciifile)
		ascii.MapReload(scanne, word)
		if len(word) != 1 && ascii.CheckNewLine(word) {
			for i := 0; i < len(word)/2; i++ {
				fmt.Println()
			}
			return
		} else {
			ascii.AsciiProcess(word)
		}
	}
}
