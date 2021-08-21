package main

import (
    "fmt"
    "time"
)

func main() {

    timer1 := time.NewTimer(5 * time.Second)

    <-timer1.C
    fmt.Println("Timer 1 fired")

    timer2 := time.NewTimer(10* time.Second)
    go func() {
        <-timer2.C
        fmt.Println("Timer 2 fired")
    }()
    time.Sleep(15 * time.Second)
}