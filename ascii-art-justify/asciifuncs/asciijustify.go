package ascii

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type jstf func(s string) string

var libOption = map[string]jstf{
	"center":  Center,
	"left":    Left,
	"right":   Right,
	"justify": Justify,
}

func JustifyContent(option, txt string) string {
	res := ""
	if justcont, ok := libOption[option]; ok {
		spc := justcont(txt)
		res = AsciiProcess(txt, spc, option)
	} else {
		fmt.Println(Danger + "Sorry we don't have this option yet\nTry another one like:\ncenter\tleft\tright\tjustify")
		os.Exit(0)
	}
	return res
}

func GetLen(word string) int {
	res := 0
	if len(word) == 0 {
		return 0
	}
	for j := 0; j < len(word); j++ {
		if word[j] < ' ' || word[j] > '~' {
			fmt.Println((Danger + "Character not found :("))
			os.Exit(0)
		}
		line := len(AsciiMap[rune(word[j])][0])
		res += line
	}
	return res
}

func GetSize() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, errcmb := cmd.Output()
	if errcmb != nil {
		log.Fatal(errcmb)
	}

	l := len(strings.Split(string(out), " ")[1])

	size, err := strconv.Atoi(strings.Split(string(out), " ")[1][:l-1])
	if err != nil {
		log.Fatal(err)
	}
	return size
}
