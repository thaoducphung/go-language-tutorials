package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

var currentPath = "E:/Study/GoProgramming"

func check(e error) {
	if e!= nil {
		panic(e)
	}
}

func main() {
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile(currentPath+"/tmp/dat1", d1, 0644)
	check(err)

	f, err := os.Create(currentPath+"/tmp/dat2")
	check(err)

	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush()
}
