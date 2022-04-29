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

	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	cErr := clipboard.Init()
	if cErr != nil {
		panic(err)
	}

	clipboard.Write(clipboard.FmtText, data)
}
