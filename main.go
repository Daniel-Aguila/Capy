package main

import (
	"fmt"
	"os"
)

func manual() {
	fmt.Println("\n      Welcome to Capyfilter\n" +
		"----parameters - explanation----\n" +
		"a -  sort directory by alphabetical order" +
		"")
	return
}

func filterAlphabetical(filesDir []os.DirEntry) {
	fmt.Println("Filtering by alphabetical order...")
}

func mapArgs(args []string) map[string]bool {
	set := make(map[string]bool)
	for _, arg := range args {
		set[arg] = true
	}
	return set
}

func main() {
	var argSet map[string]bool

	argSet = mapArgs(os.Args[1:])
	if argSet["-h"] || argSet["--help"] {
		manual()
	}
}
