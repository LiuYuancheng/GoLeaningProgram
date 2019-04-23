package main

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
)

type Robot interface {
	PowerON() error
}

func Boot(r Robot) error{
	return r.PowerON()
}

type T800 struct{
	Name string
}

func (a *T800) PowerON() error{
	return nil
}

type R2D2 struct {
	Broken bool
}

func (r *R2D2) PowerON() error{
	if r.Broken {
		return errors.New("R2D2 is broken")
	} else{
		return nil
	}
}

func main(){
	t := T800{
		Name: "The Terminator",
	}

	r := R2D2{
		Broken:	true,
	}
	err := Boot(&r)


	if err != nil{
		fmt.Println(err)
	} else{
		fmt.Println("Robot is powered on.")
	}

	err = Boot(&t)
	if err != nil{
		fmt.Println(err)
	} else{
		fmt.Println("Robot is powered on.")
	}
	fmt.Println("Byte buffer test")
	var buffer bytes.Buffer
	for i:=0; i<500; i++{
		buffer.WriteString("Z")
	}
	fmt.Println(strings.ToLower(buffer.String()))
	// string searching test
	testSting := "    This is  the test string"
	//testSting := "123"
	fmt.Println(strings.Index(testSting, "this"))
	fmt.Println(strings.Index(testSting, "notAppear"))
	fmt.Println(strings.TrimSpace(testSting))

	fmt.Println(strings.Replace(testSting,"This","that", -1 ))





}
