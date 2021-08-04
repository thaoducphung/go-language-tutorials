package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		// handle the error here
		fmt.Println(err)
		return
	}
	defer file.Close()

	// get the file size
	stat, err := file.Stat()
	if err != nil {
		// handle errors
		fmt.Println(err)
		return
	}
	// read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		// handle error 
		fmt.Println(err)
		return
	}

	str := string(bs)
	fmt.Println(str)
}
