package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	bs, err := ioutil.ReadFile("test.txt")
	if err != nil {
	 // handle error
	 	return
	}

	str := string(bs)
	fmt.Println(str)
}
