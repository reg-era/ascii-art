package main

import (
	"bufio"
	"fmt"
	"os"

	ascii "ascii/asciifuncs"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Too many argument :)")
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
	ascii.AsciiCorrect(word)
}
