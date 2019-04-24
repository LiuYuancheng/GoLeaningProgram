package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"io"
	"io/ioutil"
	"log"
	"os"
)

type Config struct{
	Name string
	Awake bool
	Hungry bool
}

func half(numberTohalf int) (int, error) {
	if (numberTohalf%2 !=0){
		return -1, fmt.Errorf("Can not do half: %v", numberTohalf)
	}
	return numberTohalf/2, nil
}

func main(){
	fileByte, err := ioutil.ReadFile("BoatVars.py")
	if err != nil{
		fmt.Println(err)
		//return
	}
	fmt.Println(fileByte)
	fmt.Println("Test writing the file")

	//b := make([]byte, 0)
	//content := "hello world"
	err = ioutil.WriteFile("IO_test.txt", fileByte, 0777)
	if err != nil{
		log.Fatal(err)
	}
	// list all the file in the dir.
	files, err := ioutil.ReadDir(".")
	if err != nil{
		log.Fatal(err)
	}
	for _, file:= range files{
		fmt.Println(file.Mode(), file.Name())
	}
	// Call the os to do the file IO.
	from, err := os.Open("./IO_test.txt")
	if err != nil{
		log.Fatal(err)
	}
	defer from.Close()

	to, err := os.OpenFile("osCreated.txt", os.O_RDWR|os.O_CREATE, 0777)
	if err != nil{
		log.Fatal(err)
	}

	_, err = io.Copy(to, from)
	if err!= nil{
		log.Fatal(err)
	}
	to.Close()

	// Test remove the file:
	fmt.Println("Test remove the IO_test.txt file")
	err = os.Remove("./osCreated.txt")
	if err != nil{
		log.Fatal(err)
	}

	// Test create a toml file
	c := Config{}
	_, err = toml.DecodeFile("config.toml", &c)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Printf("%+v", c)
	}
