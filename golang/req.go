package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func TrimFirstAndLast(s string) string {
	if len(s) > 44 {
		s = s[43 : len(s)-1]
	}
	return s
}

func main() {
	//resp, err := http.Get("http://192.168.1.103:8000/auth?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJsb2dpbiI6IjEyMzQ1Njc4OTAiLCJwYXNzd29yZCI6IkpvaG4gRG9lIn0.qupq7x2Xp32sZbOk9wH49EefoYyiuyYEd8sG1rs_lfA")

	resp, err := http.Get("http://192.168.1.103:8000/getall?type=alc")
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	log.Printf(sb)
	/*
		str := TrimFirstAndLast(sb)
		log.Println(str)*/
}

//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJsb2dpbiI6IjEyMzQ1Njc4OTAiLCJwYXNzd29yZCI6IkpvaG4gRG9lIn0.qupq7x2Xp32sZbOk9wH49EefoYyiuyYEd8sG1rs_lfA
