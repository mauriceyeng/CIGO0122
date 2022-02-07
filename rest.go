package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Response struct {
	//attributes
	camera float32
	color  string
	ram    int
	rom    int
	price  float32
	size   float32
}

func main() {
	url := "https://reqres.in/api/products/3"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(response.Body)
	var responseObject Response
	json.Unmarshal(data, &responseObject)
	log.Println(responseObject)
}
