package ascii

import (
	"fmt"
	"os"
)

var libColor = map[string]string{
	"red":     "\033[31m",
	"green":   "\033[32m",
	"yellow":  "\033[33m",
	"blue":    "\033[34m",
	"purple":  "\033[35m",
	"bluesky": "\033[36m",
	"gray":    "\033[37m",
	"white":   "\033[97m",
}

func SearchColor(color string) string {
	res := ""
	if FindColor, ok := libColor[color]; ok {
		res = FindColor
	} else {
		fmt.Println(Danger + "Sorry we don't have this color yet\nTry another one like:\n\033[31mred\t\033[32mgreen\t\033[33myellow\t\033[34mblue\t\033[36mbluesky\t\033[35mpurple\t\033[37mgray\t\033[97mwhite")
		os.Exit(0)
	}
	return res
}
