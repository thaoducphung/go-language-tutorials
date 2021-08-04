package main

import "os"

func main() {
	file, err := os.Create("test1.txt")
	if err!= nil {
		// handle error here
		return
	}
	defer file.Close()

	file.WriteString("test")
}
