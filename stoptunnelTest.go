package main

import (
	"fmt"
	"time"
)

func sender(message chan string){
	t := time.NewTicker(1*time.Second)
	for{
		message	<- "send a sample message"
		<-t.C
	}
}

func main(){
	message := make(chan string)
	stop := make(chan bool)
	go sender(message)
	go func(){
		time.Sleep(time.Second*2)
		fmt.Println("Time is up")
		stop <- true
	}()
	for{
		select{
			case <-stop:
				return
			case msg:= <- message:
				fmt.Println(msg)
		}
	}
}