package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"golang.design/x/clipboard"
)

const NO_FILENAME = 1

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "No filename")
		os.Exit(NO_FILENAME)
	}

	file := os.Args[1]

	data, rErr := ioutil.ReadFile(file)
	if rErr != nil {
		panic(rErr)
	}

	cErr := clipboard.Init()
	if cErr != nil {
		panic(cErr)
	}

	clipboard.Write(clipboard.FmtText, data)
}
