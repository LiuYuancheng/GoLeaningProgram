package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name string
	Age int
	Hobbies []string
}

type PersonJson struct{
	Name string		"json:\"name, omitempty\""
	Age int			"json:\"age, omitempty\""
	Hobbies [] string "json:\"hobbies, omitempty\""
}

func main(){
	hobbies := []string{"Cycling", "Cheese", "Techno"}
	p := PersonJson{
		Name:	"George",
		Age: 40,
		Hobbies: hobbies,
	}
	//p = PersonJson{}
	fmt.Printf("<+%v>\n", p)
	jsonByteData, err := json.Marshal(p)
	if err != nil{
		log.Fatal(err)
	}
	jsonString := string(jsonByteData)
	fmt.Println(jsonString)
	// decode the json data to struct.
	jsonStringData := "{\"name\":\"George\",\"age\":40,\"hobbies\":[\"Cycling\",\"Cheese\",\"Techno\"]}"
	jsonByteData = []byte(jsonStringData)
	p = PersonJson{}
	err = json.Unmarshal(jsonByteData, &p)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", p)
}



