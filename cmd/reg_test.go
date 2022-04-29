package cmd

import (
	"encoding/base64"
	"fmt"
	"log"
	"regexp"
	"testing"
)

var (
	re = regexp.MustCompile(`\!\[.*\]\((.*)\).*`)
	reImgTagStr = regexp.MustCompile(`<img.+?src='(/attachments.+?)'.*?>`)
)

func TestReg(t *testing.T) {

	mdContent :=`
	これに沿えばOK
https://dev.classmethod.jp/articles/gcp-billing-quota-increase/

<img title='スクリーンショット 2021-03-08 10.58.17.png' alt='スクリーンショット 2021-03-08 10.58.17' src='/attachments/77bee1af-1195-4e47-ae7f-c0ad4f86febd' width="1212" data-meta='{"width":1212,"height":1446}'>

https://cloud.google.com/billing/docs/how-to/manual-payment?hl=ja&visit_id=637168173110950031-955002622&rd=1

	英語のつたなさ・・・

	<img title='スクリーンショット 2021-04-06 14.44.03.png' alt='スクリーンショット 2021-04-06 14.44.03' src='/attachments/a6bd9f66-2839-4d9d-a108-db572b7897ce' width="1554" data-meta='{"width":1554,"height":1516}'>
`

	urls := reImgTagStr.FindAllStringSubmatch(mdContent, -1)

	for _, str := range urls {
		if len(str) >= 2 {
			url := str[1]
			log.Println(url)
		}
	}
}

func TestID(t *testing.T) {
	input := 5380

	inputWithBlogSuffix := fmt.Sprintf("Blog/%d", input)

	s := base64.StdEncoding.EncodeToString([]byte(inputWithBlogSuffix))

	//expect := "QmxvZy81Mzgw"

	t.Logf("s is %s", s)
}