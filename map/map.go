package main

import "fmt"

func main() {
	x := make(map[string]int)
	x["key"] =10
	fmt.Println(x)
	fmt.Println(x["key"])
	name, ok := x["hello"]
	fmt.Println(name,ok)
	if name, ok = x["hey"]; ok {
		fmt.Print(name,ok)
	}
}
