package main

import (
	"fmt"
	"os"
)

func manual() {
	fmt.Println("\n      Welcome to Capyfilter\n" +
		"----parameters - explanation----\n" +
		"a -  sort directory by alphabetical order" +
		"" +
		"\n")
}

func filterAlphabetical(filesDir []os.DirEntry) {
	fmt.Println("Filtering by alphabetical order...")
}

func main() {
	fmt.Println(os.Args)
	if os.Args[1] == "-h" || os.Args[1] == "--help" {
		manual()
	}
}
