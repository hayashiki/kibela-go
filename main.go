package main

import (
	"log"

	"github.com/hayashiki/kibela-go/kibela"
)

func main() {
	client, err := kibela.NewClient(nil, "hayashiki")
	if err != nil {
		log.Fatal(err)
		return
	}
	resp, err := client.Note.Create()
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Printf("log %v", resp)
}
