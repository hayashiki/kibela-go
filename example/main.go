package main

import (
	"log"
	"os"

	"github.com/hayashiki/kibela-go"
)

func main() {
	team := os.Getenv("KIBELA_TEAM")
	accessToken := os.Getenv("KIBELA_TOKEN")

	client, err := kibela.NewClient(nil, team, accessToken)
	if err != nil {
		log.Fatal(err)
		return
	}
	notes, err := client.Note.Search(os.Args[1])
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, note := range notes {
		log.Printf("log %v", note.Title)
	}
}
