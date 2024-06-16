package main

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"io/ioutil"
	"log"
	"net/http"
)

type prods1 struct {
	Name      string      `json:"name"`
	Daystitle string      `json:"daystitle"`
	Shops_ids []string    `json:"shops_Ids"`
	Imagefull interface{} `json:"imagefull"`
}

type prodlist1 struct {
	Products []prods1 `json:"products"`
}

func main() {
	resp, err := http.Get("http://192.168.1.100:6969/getall?type=desert")
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	var data bson.M
	//log.Println(string(body))
	sb := string(body)
	//log.Println(sb)

	log.Println(data)
}
