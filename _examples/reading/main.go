package main

import (
	"fmt"

	"github.com/gocoderio/easyfs"
)

func main() {
	fs := easyfs.FS
	err := fs.AddFile("hello.txt", "Hello, world!")
	if err != nil {
		panic("Error adding file to filesystem: " + err.Error())
	}

	data, err := fs.ReadFile("hello.txt")
	if err != nil {
		panic("Error reading file: " + err.Error())
	}
	fmt.Printf("data==%v\n", string(data))
}
