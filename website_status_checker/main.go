package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := [5]string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	for _, link := range links {
		checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		fmt.Print("here's the error:", err)
		return
	}
	fmt.Println(link, "is up!")
}
