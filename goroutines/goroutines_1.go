package main

import "fmt"

func f(num int) {
	for i := 0; i < 10; i ++ {
		fmt.Println(num, ":", i)
	}
}
func main() {
	for i := 0; i < 10; i ++ {
		go f(1)
	}
	var input string
	fmt.Scanln(&input)
}
