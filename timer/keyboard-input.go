package main

import (
	"time"
	"fmt"
	"github.com/micmonay/keybd_event"
)

func main() {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}
// Select keys to be pressed
	kb.SetKeys(keybd_event.VK_Q) 
	fmt.Println(keybd_event.VK_Q)
	fmt.Printf("%T\n", keybd_event.VK_Q)
// Set shift to be pressed
	kb.HasALT(true)

// Or you can use Press and Release
	kb.Press()
	time.Sleep(10 * time.Millisecond)
	kb.Release()
}
