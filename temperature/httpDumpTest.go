package temperature

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"time"
)

func main(){
	//debug := "1"
	debug := os.Getenv("DEBUG")
	client := &http.Client{
		Timeout: 50*time.Second,
	}
	request, err := http.NewRequest("GET", "https://ifconfig.io/", nil)
	if err != nil{
		log.Fatal(err)
	}
	if debug == "1"{
		debugRequest, err := httputil.DumpRequestOut(request, true)
		if err != nil{
			log.Fatal(err)
		}
		fmt.Printf("DumpRequest out\n %s", debugRequest)
	}
	response, err := client.Do(request)
	defer response.Body.Close()


	if debug == "1"{
		debugResponse, err := httputil.DumpResponse(response, true)
		if err != nil{
			log.Fatal(err)
		}
		fmt.Printf("DumpResponse:\n %s ", debugResponse)
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Printf("%s", body)
}