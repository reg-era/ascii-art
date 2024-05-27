package main

import (
	ascii "ascii/pkg/Ascii"
	outil "ascii/pkg/Outil"
	"os"
)

func main() {
	Outils := outil.Outils{
		Args: os.Args[1:],
	}
	ascii.Process_Text(&Outils)
}
