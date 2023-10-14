package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

// func

func main() {
	var data []byte
	var err error
	fi, _ := os.Stdin.Stat()
	if (fi.Mode() & os.ModeCharDevice) == 0 {
		data, err = io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println(err)
			return
		}

	}

	cFlag := flag.Bool("c", false, "count bytes in files")
	lFlag := flag.Bool("l", false, "no of lines in files")
	wFlag := flag.Bool("w", false, "no of words in files")
	mFlag := flag.Bool("m", false, "no of characters in files")
	flag.Parse()
	var filePath string
	if len(data) == 0 {
		filePath = flag.Args()[0]
		data, err = os.ReadFile(filePath)
		if err != nil {
			fmt.Println(err)
			return
		}

	}

	text := strings.TrimSpace(string(data))
	isOneOfFlagSet := false
	if *cFlag {
		fmt.Printf("%v ", countBytes(text))
		isOneOfFlagSet = true
	}

	if *lFlag {
		fmt.Printf("%v ", countLines(text))
		isOneOfFlagSet = true
	}

	if *wFlag {
		fmt.Printf("%v ", countWords(text))
		isOneOfFlagSet = true
	}

	if *mFlag {
		fmt.Printf("%v ", countCharacters(text))
		isOneOfFlagSet = true
	}

	if isOneOfFlagSet {
		fmt.Printf("%s", filePath)
	} else {
		fmt.Printf("%v %v %v %v %s", countBytes(text), countLines(text),countWords(text),countCharacters(text),filePath)
	}
	fmt.Println()



}

func countCharacters(text string) int {
	return utf8.RuneCountInString(text)
}

func countWords(text string) int {
	return len(strings.Fields(text))
}

func countLines(text string) int {
	return len(strings.Split(text, "\n"))
}

func countBytes(text string) int {
	return len([]byte(text))
}
