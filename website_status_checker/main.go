package main

import (
	"fmt"
	"net/http"
)

// ^^^main routine created when we launched our program
func main() {
	links := [5]string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {

		// going sequentially (waiting status one-by-one) is much more slower
		// checkLink(link)

		// going concurrently with go routine
		// ^^^child routines created by the `go` keyword
		go checkLink(link, c)

		// !concurrency is not parallelism!
		// scheduler runs one routine until it finished or makse a blocking call (like an HTTP request)

		// *concurrency* - we can have multiple threads executing code.
		// if one thread blocks, another one is picke dup and worked on
		// *parallelism* - multiple threads executed at the exact same time.
		// this requires multiple CPU's.
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		fmt.Print("here's the error:", err)
		c <- "Might be down I think"
		return
	}
	fmt.Println(link, "is up!")
	c <- "Yep it's up"
}
