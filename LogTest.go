package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Animal struct {
	Name string
	Color string
}



func main() {
	testCase := 0
	a := Animal{
		Name : "cat",
		Color: "black",
	}
	fmt.Printf("%v \n", a)
	fmt.Printf("%+v \n", a)

	if testCase == 1 {
		f, err := os.OpenFile("test.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		log.SetOutput(f)
		for i := 1; i < 300; i++ {
			log.Printf("Log iterate <%d> times \n", i)
		}
	}else if testCase ==2 {
		for i:=1; i<100; i++{
			log.Printf("Log interation %d", i)
		}
	}else{
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Guess the name of my pet to win a price: ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		if text == "John"{
			fmt.Println("You have successful get the name")
		} else{
			fmt.Println("The input is not currect")
		}
	}
}




