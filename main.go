package main

import (
	"fmt"
	//tea "github.com/charmbracelet/bubbletea"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	method := os.Args[1]
	if method == "get" {
		url := os.Args[2]

		if !strings.HasPrefix(url, "http") {
			url = "http://" + url
		}
		res, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}

		defer res.Body.Close()
		fmt.Println(res.Status)
	}
}
