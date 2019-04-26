package main

import (
	"fmt"
	"github.com/golang/example/stringutil"
)

type SuperHero struct{
	Name string
	age int
	address Address
}

type Address struct {
	num int
	street string
	city string
}

func main()  {

	e:= SuperHero{
		Name: "BatMan",
		age: 32,
		address: Address{
			num: 123345,
			street: "Mountain Drive",
			city: "Gotham",
		},
	}

	fmt.Printf("%+v", e)
	fmt.Println(e.address.street)
	fmt.Println(stringutil.Reverse("testing"))



	defer fmt.Println("I am the first defer statement")
	defer fmt.Println("I am the second defer statement")
	defer fmt.Println("I am the third defer statement")

	fmt.Println("Hello world")
}

