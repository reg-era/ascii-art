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
	"center": Center,
	"left":   Left,
	"right":  Right,
	//"justify": Center,
}

func JustifyContent(option, txt string) string {
	res := ""
	if justcont, ok := libOption[option]; ok {
		spc := justcont(txt)
		res = AsciiProcess(txt, spc)
	} else {
		fmt.Println(Danger + "Sorry we don't have this option yet\nTry another one like:\ncenter\tleft\tright\tjustify")
		os.Exit(0)
	}
	return res
}

func Center(s string) string {
	res := ""
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
	spc := (size / 2) - (len(s) * 2)
	for spc != 0 {
		res += " "
		spc--
	}
	return res
}

func Left(s string) string {
	res := ""
	return res
}

func Right(s string) string {
	res := ""
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
	spc := (size) - (len(s) * 8)
	for spc != 0 {
		res += " "
		spc--
	}
	return res
}
