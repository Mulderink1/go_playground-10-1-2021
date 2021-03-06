package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string {
		"https://www.google.com/",
		"https://www.facebook.com/",
		"https://stackoverflow.com/",
		"https://golang.org/",
		"https://www.amazon.com/",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go func(l string) {
			time.Sleep(5 * time.Second)
			checkLink(l, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		c <- link + " might be down!"
		return
	}

	fmt.Println(link)
	c <- link
}