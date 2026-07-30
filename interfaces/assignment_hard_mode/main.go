package main

import (
	"fmt"
	"io"
	"os"
)

type logWriter struct {
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Print("\t~Context of the file:\n")
	fmt.Println(string(bs))
	fmt.Println("\t~Just logged this many bytes:", len(bs))
	return len(bs), nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("There are not enough arguments!")
		fmt.Println("Missing the filename as the second argument!")
		os.Exit(1)
	}

	filename := os.Args[1]
	fmt.Println("filename is:", filename)

	// https://golang.org/pkg/os/#Open
	file, error := os.Open(filename)
	if error != nil {
		fmt.Println("During attempt to open file", filename, "error occured")
		fmt.Println(error)
		os.Exit(1)
	}

	io.Copy(logWriter{}, file)
}
