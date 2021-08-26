package main

// Reference
// https://github.com/micmonay/keybd_event

import (
	"time"
	"fmt"
	"github.com/micmonay/keybd_event"
)

const VARIANT_TIME = 1
const HUNT_TIME int = 60 + VARIANT_TIME
const HEALING_TIME int = 60 * 60 + VARIANT_TIME
const TRAINING_TIME int = 60 * 15 + VARIANT_TIME
const FARMING_TIME int = 60 * 10 + VARIANT_TIME
const CHOPPING_TIME int = 60 * 5 + VARIANT_TIME
const ADVENTURE_TIME int = 60* 60 + VARIANT_TIME

const HUNT_IMAGE string = "E:/Study/GoProgramming/timer/Dagger.png"
const CHOP_IMAGE string = "E:/Study/GoProgramming/timer/Axe.png"
const FRAMING_IMAGE string = "E:/Study/GoProgramming/timer/Carrot.png"
const TRAINING_IMAGE string = "E:/Study/GoProgramming/timer/Star.png"


var HUNT_WAIT_TIME     = 0          // second

var HEALING_WAIT_TIME = 12 + 0 * 60 // second

var TRAINING_WAIT_TIME = 8 + 10 * 60 // second

var FARMING_WAIT_TIME  = 4 + 0 * 60 // second

var ADVENTURE_WAIT_TIME  = 10 + 40 * 60 // second

var CHOPPING_WAIT_TIME = 2 + 1 * 60 // second
// var CHOPPING_WAIT_TIME = 0 // second


func runKeyboard(keyboard_input int) {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}
// Select keys to be pressed
	kb.SetKeys(keyboard_input) 
	
// Set shift to be pressed
	kb.HasALT(true)

// Or you can use Press and Release
	kb.Press()
	time.Sleep(10 * time.Millisecond)
	kb.Release()
}

func bgTask(time_ticker int, time_wait int, title string, keyboard_input int) {
	
	time.Sleep(time.Duration(time_wait) * time.Second)
	runKeyboard(keyboard_input)
	fmt.Println(title," Time")
	
	ticker := time.NewTicker(time.Duration(time_ticker) * time.Second)
	for _ = range ticker.C {
		fmt.Println(title," Time")
		runKeyboard(keyboard_input)
	}
}

func main() {
	// notify()
	
	go bgTask(HUNT_TIME, HUNT_WAIT_TIME, "Hunt", keybd_event.VK_Q)
	
	go bgTask(CHOPPING_TIME, CHOPPING_WAIT_TIME, "Axe" , keybd_event.VK_W)
	
	// go bgTask(HEALING_TIME, HEALING_WAIT_TIME, "Heal" , keybd_event.VK_H)
	
	// go bgTask(FARMING_TIME, FARMING_WAIT_TIME, "Farming" , keybd_event.VK_F)
	
	go bgTask(ADVENTURE_TIME, ADVENTURE_WAIT_TIME, "Adventure" , keybd_event.VK_A)
	
	go bgTask(TRAINING_TIME, TRAINING_WAIT_TIME, "Training" , keybd_event.VK_E)
	
	select {}
}

