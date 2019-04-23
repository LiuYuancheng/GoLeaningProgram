package main

import (
	"fmt"
	"time"
)

func channelReader(message <- chan string){
	msg:= <- message
	fmt.Println(msg)
}

func channelWriter(message chan <- string, msg string){
	message <- msg
}

func channelReadAndWriter(message chan string, Inmsg string){
	if Inmsg != nil{
		msg := <- message
		fmt.Println(msg)
	}else{
		message <- Inmsg
	}
}


func pinger (c chan string){
	t := time.NewTicker(time.Second *1)
	count := 0
	fmt.Println("Enter here")
	for {
		count += 1
		c <- "ping"
		fmt.Println(count)
		<-t.C // return from here.
	}
	fmt.Println("Leave the function")
}

func main(){
	message := make(chan string)
	go pinger(message)
	for {
		msg := <-message
		fmt.Println(msg)
	}
	}