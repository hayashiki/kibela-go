package main

import (
	"bytes"
	"encoding/base64"
	"image/png"
	"log"
	"os"
	"strings"

	"github.com/hayashiki/kibela-go"
)

var (
	team = os.Getenv("KIBELA_TEAM")
	accessToken = os.Getenv("KIBELA_TOKEN")
	client *kibela.Client
)

func init() {
	k, err := kibela.NewClient(nil, team, accessToken)
	if err != nil {
		log.Fatal(err)
		return
	}
	client = k
}

func main() {
	attachment()
	return

	note()
	return

	notes, err := client.Note.Search(os.Args[1])
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, note := range notes {
		log.Printf("log %v", note.Title)
	}
}

func note() {
	note, err := client.Note.Get("QmxvZy81Mzgw")
	if err != nil {
		log.Printf("log %v", err)
	}
	log.Printf("note %v", note)
}

func attachment() {
	att, _ := client.Attachment.Get("/attachments/a6bd9f66-2839-4d9d-a108-db572b7897ce")
	log.Printf("att.DataUrl %v", att.DataUrl)
	dataurl := strings.Replace(att.DataUrl, "data:image/png;base64,", "", 1)

	//f, err := os.Create("example2.png")
	f, err := os.OpenFile("example.png", os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
	}

	d, err := base64.StdEncoding.DecodeString(dataurl)
	r := bytes.NewReader(d)
	img, err := png.Decode(r)

	//    reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(dataurl))
	//    c, _, err := image.DecodeConfig(reader)
	//    if err != nil {
	//        log.Fatal(err)
	//    }

	//	mimeType := http.DetectContentType(bytes)
	//
	//	switch mimeType {
	//	case "image/jpeg":
	//		base64Encoding += "data:image/jpeg;base64,"
	//	case "image/png":
	//		base64Encoding += "data:image/png;base64,"
	//	}

	png.Encode(f, img)

	//f.Write(img)
	//
	//f2, err := os.OpenFile("example2.png", os.O_WRONLY|os.O_CREATE, 0777)
	//if _, err := io.Copy(f2, r); err != nil {
	//	return
	//}
}