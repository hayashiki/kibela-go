package main

import (
	"log"

	"github.com/hayashiki/kibela-go"
)

func main() {
	client, err := kibela.NewClient(nil, "hayashiki")
	if err != nil {
		log.Fatal(err)
		return
	}
	resp, err := client.Note.Search("aaa")
	// resp, err := client.Note.GetAll()
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Printf("log %v", resp)
}
