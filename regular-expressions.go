package main

import (
	"bytes"
	"fmt"
	re "regexp"
)

func main() {
	match, _ := re.MatchString("p([a-z]+)ch","peach")
	fmt.Println(match)

	r, _ := re.Compile("p([a-z]+)ch")

	fmt.Println("1", r.MatchString("peach"))

	fmt.Println("2", r.FindString("peach punch"))

	fmt.Println("3", r.FindStringIndex("peach punch"))

	fmt.Println("4", r.FindStringSubmatch("peach punch"))

	fmt.Println("5", r.FindStringSubmatchIndex("peach punch"))

	fmt.Println("6", r.FindAllStringSubmatchIndex("peach punch pinch", -1))

	fmt.Println("7", r.FindAllString("peach punch pinch", 2))

	fmt.Println("8", r.Match([]byte("peach")))

	r = re.MustCompile("p([a-z]+)ch")
	fmt.Println("9", r)

	fmt.Println("10", r.ReplaceAllString("a peach","<fruit>"))

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println("11", string(out))
}
