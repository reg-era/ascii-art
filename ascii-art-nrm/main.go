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
		fmt.Println(Danger + "Please insert text you want to print :)")
		os.Exit(0)
	}
	if len(args) > 1 {
		fmt.Println(Danger + "Too many argument :)")
		os.Exit(0)
	}
	word := args[0]
	asciifile, err := os.Open("asciilib/standard.txt")
	if err != nil {
		fmt.Println("Dont play in program file >(")
		os.Exit(0)
	}
	defer asciifile.Close()
	scanne := bufio.NewScanner(asciifile)
	ascii.MapReload(scanne)
	if len(word) != 1 && ascii.CheckNewLine(word) {
		for i := 0; i < len(word)/2; i++ {
			fmt.Println()
		}
		return
	} else {
		ascii.AsciiProcess(word)
	}
}
