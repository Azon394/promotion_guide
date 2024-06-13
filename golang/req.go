package main

import (
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://192.168.1.103:8000/auth?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJsb2dpbiI6IjEyMzQ1Njc4OTAiLCJwYXNzd29yZCI6IkpvaG4gRG9lIn0.qupq7x2Xp32sZbOk9wH49EefoYyiuyYEd8sG1rs_lfA")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(resp)
}

//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJsb2dpbiI6IjEyMzQ1Njc4OTAiLCJwYXNzd29yZCI6IkpvaG4gRG9lIn0.qupq7x2Xp32sZbOk9wH49EefoYyiuyYEd8sG1rs_lfA
