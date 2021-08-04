package main

import (
	"fmt"
	"os"
)

func main() {
	dir, err := os.Open(".")
	if err != nil {
		// handle error
		fmt.Println(err)
		return
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		// handle error
		fmt.Println(err)
		return
	}
	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}

}
