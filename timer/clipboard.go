package main

import (
	"fmt"
	"time"
	"github.com/atotto/clipboard"
)

func main() {
	clipboard.WriteAll("rpg hunt")
	text, _ := clipboard.ReadAll()
	fmt.Println(text)
	time.Sleep(10 * time.Second)
	clipboard.WriteAll("rpg axe")
	text, _ = clipboard.ReadAll()
	fmt.Println(text)
	
}