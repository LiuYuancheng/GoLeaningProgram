package main
// This program is used to test the time tick.

import (
	"fmt"
	"log"
	"time"
)

func tickerFunc(){
	fmt.Println("Enter the ticker test Func. ")
	c := time.Tick(1*time.Second)
	for t := range c{
		fmt.Printf("The time is now %v \n", t)
	}
}

func UTCtoMST(utcTime string) string{
	// Convert the UTC time to MST time .
	return "2006-01-02T15:04:05+00:00"
}

// Question here: how to Parse UTC time.
func main(){
	fmt.Println("Test the time.Parse function: ")
	tString := time.Now().String()
	fmt.Println(tString)
	tString = "2006-01-02T15:04:05+00:00"
	t, err := time.Parse(time.RFC3339, tString)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(t.Date())
	time.Sleep(3*time.Second)
	fmt.Println("I am awake up")
	go tickerFunc()
	for{
		select{
			case <- time.After(10 * time.Second):
				fmt.Println("Time is up in 2 second.")
				return
		}
	}
	// test to use the tick
}