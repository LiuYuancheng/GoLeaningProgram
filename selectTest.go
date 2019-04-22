package main

import (
	"fmt"
	"time"
)

func ping1(c chan string){
	time.Sleep(time.Second * 1)
	c <- "Ping on channel 1"
}

func ping2(c chan string){
	time.Sleep(time.Second * 2)
	c <- "Ping on channel 2"
}

func main(){
	channel1 := make(chan string)
	channel2 := make(chan string)

	go ping1(channel1)
	go ping2(channel2)

	select{
		case msg1 := <- channel1:
			fmt.Println(msg1)
		case msg2 := <- channel2:
			fmt.Println(msg2)
		// Add the timeout part
		case <- time.After(500 * time.Millisecond):
			fmt.Print("No message received, giving up")
	}
	time.Sleep(time.Second *3)
}