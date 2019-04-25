package main

import (
	"fmt"
	"strconv"
)

type Movie struct {
	Name string
	Rating float32
}

func main(){
	var array [2] int
	array[0] = 1
	array[1] = 2
	fmt.Println(array)

	var slice = make([] string,2)
	slice[0] = "element 1"
	slice[1] = "element 2"
	fmt.Println("Append a element in the slice")
	slice = append(slice, strconv.Itoa(array[0]))
	slice = append(slice[:1], slice[1+1:]...)
	fmt.Println(slice)

	var players = make(map[string] int)
	players["p1"] = 1
	players["p2"] = 2

	m := Movie{Name: "Citizen Kane",
		Rating: 10,
	}

	fmt.Println(m)
	
}