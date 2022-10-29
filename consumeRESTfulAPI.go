//https://www.youtube.com/watch?v=bYZlQMWR8HM

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	response, err := http.Get("https://api.coinbase.com/v2/prices/spot?currency=USD")

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s \n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
	jsonData := map[string]string{"firstname": "NIC", "lastname": "Raboy"}
	jasonValue, _ := json.Marshal(jsonData)
	request, _ := http.NewRequest("POST", "http://httpbin.org/post", bytes.NewBuffer((jasonValue)))
	request.Header.Set("Content-type", "application/json")
	client := &http.Client{}
	response, err = client.Do(request)
	//response,err = http.Post("http://httpbin.org/post", "application/json", bytes.NewBuffer(jasonValue))
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s \n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}
