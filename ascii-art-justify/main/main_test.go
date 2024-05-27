package main

import (
	"fmt"
	"os"
	"testing"

	Ascii "ascii/pkg/Ascii"
	Outil "ascii/pkg/Outil"
)

type AsciiTestCase struct {
	args         []string // Arguments to pass to Ascii.Process_Text
	expectedFile string   // Expected file name to read the output from
}

func TestProcess_Text(t *testing.T) {
	// Define test cases as a slice of AsciiTestCase
	testCases := []AsciiTestCase{
		{args: []string{"hello"}, expectedFile: "test01.txt"},
		{args: []string{"12~/*AB"}, expectedFile: "test02.txt"},
		{args: []string{"{T\\'/;+543}", "thinkertoy"}, expectedFile: "test03.txt"},
		{args: []string{"--align=left", "hello word"}, expectedFile: "test04.txt"},
		{args: []string{"--align=right", "hello word"}, expectedFile: "test05.txt"},
		{args: []string{"--align=center", "hello word"}, expectedFile: "test06.txt"},
	}

	for _, testCase := range testCases {
		os.Setenv("COLUMNS", "150")
		outil := &Outil.Outils{Args: testCase.args}
		Ascii.Process_Text(outil)
		expectedOutput := ReadFile(testCase.expectedFile)
		if outil.Result != expectedOutput {
			t.Errorf("Test case failed. Expected:\n%s\nActual:\n%s", expectedOutput, outil.Result)
		} else {
			fmt.Println("Test valid")
		}
	}
}

func ReadFile(file string) string {
	Data, _ := os.ReadFile("test/" + file)
	return string(Data)
}
