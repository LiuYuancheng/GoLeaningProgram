package main

import (
	"fmt"
	"time"
)

func slowFunc(m chan string){
	time.Sleep(time.Second*2)
	m <- "flowFunc() finished"
	fmt.Println("Sleeper() finished")
}

func bufRelease(c chan string ){
	// release items from the buffer tunnel
	for msg:= range c {
		fmt.Println(msg)
	}
}


func main(){
	m := make(chan string)
	go slowFunc(m)

	msg := <-m
	fmt.Println(msg)

	fmt.Println("Now show here before sleeper()")
	time.Sleep(time.Second*3)

	message := make(chan string, 2)
	message <- "hello world 1"
	message <- "hello world 2"
	close(message)
	time.Sleep(time.Second *1)
	bufRelease(message)
	}