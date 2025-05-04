package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func getAPI(url string) {

	if !strings.HasPrefix(url, "http") {
		url = "http://" + url
	}
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	header := res.Header.Get("Content-Type")
	if strings.Contains(header, "application/json") {

	}

	defer res.Body.Close()
	fmt.Println(res.StatusCode)
}

func main() {
	method := os.Args[1]
	if method == "get" {
		url := os.Args[2]
		getAPI(url)
	}
}
