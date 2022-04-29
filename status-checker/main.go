package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	pages := []string{
		"https://google.com",
		"https://facebook.com",
		"https://amazon.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://draw.io",
	}

	var results = make(map[string]string)
	channel := make(chan string)

	for _, link := range pages {
		go checkLink(link, channel)
	}

	fmt.Println(results)
	fmt.Println(<-channel)
}
func checkLink(link string, c chan string) {
	respone, err := http.Get(link)
	if err != nil {
		os.Exit(1)
	}
	fmt.Println("Status of", link, "is", respone.Status)

	// send message to channel
	c <- respone.Status
}
