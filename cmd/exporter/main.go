package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/hayashiki/kibela-go"
	"golang.org/x/sync/errgroup"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)

var (
	team = os.Getenv("KIBELA_TEAM")
	accessToken = os.Getenv("KIBELA_TOKEN")
	client *kibela.Client
)

var (
	re = regexp.MustCompile(`<img.+?src='(/attachments.+?)'.*?>`)
)

func init() {
	k, err := kibela.NewClient(nil, team, accessToken)
	if err != nil {
		log.Fatal(err)
		return
	}
	client = k
}

func main()  {

	//入力値 4951
	num := 4951
	// Blog/5380にしてから、エンコードしてIDにする
	//QmxvZy81Mzgw
	blogId := fmt.Sprintf("Blog/%d", num)
	id := base64.StdEncoding.EncodeToString([]byte(blogId))
	note, err := client.Note.Get(id)
	if err != nil {
		return
	}

	mdContent := note.Content

//			`
//これに沿えばOK
//https://dev.classmethod.jp/articles/gcp-billing-quota-increase/
//
//<img title='スクリーンショット 2021-03-08 10.58.17.png' alt='スクリーンショット 2021-03-08 10.58.17' src='/attachments/77bee1af-1195-4e47-ae7f-c0ad4f86febd' width="1212" data-meta='{"width":1212,"height":1446}'>
//
//https://cloud.google.com/billing/docs/how-to/manual-payment?hl=ja&visit_id=637168173110950031-955002622&rd=1
//
//英語のつたなさ・・・
//
//<img title='スクリーンショット 2021-04-06 14.44.03.png' alt='スクリーンショット 2021-04-06 14.44.03' src='/attachments/a6bd9f66-2839-4d9d-a108-db572b7897ce' width="1554" data-meta='{"width":1554,"height":1516}'>
//`

	//	mdからpath　をしゅとくする
	var attachments []string
	urls := re.FindAllStringSubmatch(mdContent, -1)

	for _, str := range urls {
		if len(str) >= 2 {
			url := str[1]
			log.Println(url)
			attachments = append(attachments, url)
		}
	}

	eg := errgroup.Group{}
	mutex := &sync.Mutex{}
	for _, attachment := range attachments {
		attachment := attachment
		eg.Go(func() error {
			att, err := client.Attachment.Get(attachment)
			if err != nil {
				log.Printf("GetMultiAsync err: %w", err)
				return err
			}
			mutex.Lock()
			dataUrl := strings.Replace(att.DataUrl, "data:image/png;base64,", "", 1)
			log.Print(att.MimeType)
			download(dataUrl, fmt.Sprintf("%s.%s", att.ID, "png"))
			//files = append(files, file)
			mutex.Unlock()
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		log.Printf("eg.Wait err: %w", err)
		//return files, err
	}
}

func download(dataUrl, path string) error {
	// Create the parent directory.
	if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
		return err
	}

	// Open the file to write to.
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create: %s", err)
	}
	defer func() { _ = f.Close() }()

	d, err := base64.StdEncoding.DecodeString(dataUrl)
	r := bytes.NewReader(d)
	img, err := png.Decode(r)
	err = png.Encode(f, img)
	return err
}
//
//// Number returns the number
//func (i ID) Number() (int, error) {
//	stuff := strings.Split(i.String(), "/")
//	if len(stuff) != 2 {
//		return 0, fmt.Errorf("invalid id: %s", string(i))
//	}
//	num, err := strconv.Atoi(stuff[1])
//	if err != nil {
//		return 0, xerrors.Errorf("invalid id: %s, error: %w", i.String(), err)
//	}
//	return num, nil
//}
//// ID represents kibela ID
//type ID string
//
//func newID(typ string, num int) ID {
//	str := fmt.Sprintf("%s/%d", typ, num)
//	return ID(base64.RawStdEncoding.EncodeToString([]byte(str)))
//}
//
//func (i ID) String() string {
//	s, _ := base64.RawStdEncoding.DecodeString(string(i))
//	return string(s)
//}
//
//// Raw returns raw id string
//func (i ID) Raw() string {
//	return string(i)
//}
//
//// Empty returns the if the id is empty
//func (i ID) Empty() bool {
//	return i.Raw() == ""
//}
//
//type AttachedImage struct {
//	url      string
//	filename string
//}
//
//func imageUrlsFromMd(mdContent string) ([]string, error) {
//	imageUrls := []string{}
//
//	re := regexp.MustCompile(`\!\[.*\]\((.*)\).*`)
//	orgImageUls := re.FindAllStringSubmatch(mdContent, -1)
//	for _, urlStr := range orgImageUls {
//		if len(urlStr) >= 2 {
//			url := urlStr[1]
//			imageUrls = append(imageUrls, url)
//		}
//	}
//	return imageUrls, nil
//}
//
//func fileNameFromUrl(url string) (string, error) {
//	re := regexp.MustCompile(`[^\/]+\/([^\/]+)$`)
//
//	fileNames := re.FindAllStringSubmatch(url, -1)
//
//	file := ""
//	for _, fileArr := range fileNames {
//		if len(fileArr) < 2 {
//
//			message := fmt.Sprintf("Failed to get filename from download URL: %s", url)
//			log.Fatal(message)
//			return "", errors.New(message)
//		}
//		file = fileArr[1]
//	}
//	if len(file) == 0 {
//		return uuid.New().String(), nil
//	}
//	return file, nil
//}
//
