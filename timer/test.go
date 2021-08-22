package main

import (
    "log"
    "gopkg.in/toast.v1"
	"time"
	"fmt"
)

const VARIANT_TIME = 4
const HUNT_TIME int = 60 + VARIANT_TIME
const TRAINING_TIME int = 60 * 15 + VARIANT_TIME
const FARMING_TIME int = 60 * 10 + VARIANT_TIME
const CHOPPING_TIME int = 60 * 5 + VARIANT_TIME

const HUNT_IMAGE string = "E:/Study/GoProgramming/timer/Dagger.png"
const CHOP_IMAGE string = "E:/Study/GoProgramming/timer/Axe.png"
const FRAMING_IMAGE string = "E:/Study/GoProgramming/timer/Carrot.png"
const TRAINING_IMAGE string = "E:/Study/GoProgramming/timer/Star.png"


var HUNT_WAIT_TIME     = 0          // second
var TRAINING_WAIT_TIME = 0 + 7 * 60 // second
var FARMING_WAIT_TIME  = 0 + 1 * 60 // second

var CHOPPING_WAIT_TIME = 5 + 1 * 60 // second
// var CHOPPING_WAIT_TIME = 0 // second


func notify(title, image_name string) {
    notification := toast.Notification{
        AppID: "Epic RPG",
        Title: title,
        Message: "",
		Icon: image_name,
		Duration: "short",
    }
    err := notification.Push()
    if err != nil {
        log.Fatalln(err)
    }
}

func bgTask(time_ticker int, time_wait int, title string, image_name string) {
	
	time.Sleep(time.Duration(time_wait) * time.Second)
	notify(title,image_name)
	
	ticker := time.NewTicker(time.Duration(time_ticker) * time.Second)
	for _ = range ticker.C {
		fmt.Println(title," Time")
		notify(title,image_name)
	}
}

func main() {
	// notify()
	
	go bgTask(HUNT_TIME, HUNT_WAIT_TIME, "Hunt", HUNT_IMAGE)
	
	go bgTask(CHOPPING_TIME, CHOPPING_WAIT_TIME, "Axe" ,CHOP_IMAGE)
	
	go bgTask(FARMING_TIME, FARMING_WAIT_TIME, "Farming" ,FRAMING_IMAGE)
	
	go bgTask(TRAINING_TIME, TRAINING_WAIT_TIME, "Training" ,TRAINING_IMAGE)
	
	select {}
}

