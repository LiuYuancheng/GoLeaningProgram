package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	pro := "GetClient"

	if pro == "GetDefault" {
		response, err := http.Get("https://ifconfig.io/")
		if err != nil {
			log.Fatal(err)
		}
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf(" Get: \n %s", string(body))
	}else if pro == "PostDefault"{
		postData := strings.NewReader("{\"some\":\"json\"}")
		response, err := http.Post("https://httpbin.org/post", "application/json", postData)
		if err != nil{
			log.Fatal(err)
		}
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		if err != nil{
			log.Fatal(err)
		}
		fmt.Printf("Post repsonse: \n %s", string(body))
	}else if pro == "GetClient"{
		client := &http.Client{}
		request, err := http.NewRequest("GET", "https://ifconfig.io/", nil)
		if err != nil{
			log.Fatal(err)
		}
		response, err := client.Do(request)
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		if err!= nil{
			log.Fatal(err)
		}
		fmt.Printf("Getclient: %s \n", body)
	}

	}