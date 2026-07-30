package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct {
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// resp (type Response) struct docs: https://golang.org/pkg/net/http/#Response
	// fmt.Println(resp)

	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	// io.Copy(os.Stdout, resp.Body)

	lw := logWriter{}
	io.Copy(lw, resp.Body)
}
