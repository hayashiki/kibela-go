package kibela

import (
	"encoding/json"
	"log"
	"strings"
)

type NoteService struct {
	client *Client
}

type Note struct {
	// ID `json:"id"`
	Title string `json"title"`
	// Content string `json"content"`
}

func (s *NoteService) Create() ([]*Note, error) {

	pa := Payload{
		Query: listNoteQuery(),
	}

	pa.Query = strings.TrimSpace(pa.Query)

	log.Printf("log pa %v", pa)

	req, err := s.client.NewRequest("POST", pa)
	if err != nil {
		return nil, err
	}

	var eres struct {
		Notes struct {
			Nodes []*Note `json:"nodes"`
		} `json:"notes"`
	}

	var res response

	// noteResp := interface{}
	resp, err := s.client.Do(req, &res)

	log.Printf("log res %v", resp)

	if err := json.Unmarshal(res.Data, &eres); err != nil {
		return nil, err
	}

	for _, note := range eres.Notes.Nodes {
		log.Printf("log res %v", note.Title)
	}

	if err != nil {
		return nil, err
	}

	// data, err := ki.cli.Do(&client.Payload{Query: listNoteQuery(num, folderID, limit > 0)})

	// return res.Notes.Nodes, err
	return nil, err
}

type Payload struct {
	Query     string      `json:"query"`
	Variables interface{} `json:"variables,omitempty"`
}

type Error struct {
	Message    string          `json:"message"`
	Locations  []ErrorLocation `json:"locations,omitempty"`
	Path       []interface{}   `json:"path,omitempty"` // string or uint
	Extensions ErrorExtensions `json:"extensions"`
}

type ErrorCode string

type ErrorLocation struct {
	Line   uint `json:"line"`
	Column uint `json:"column"`
}

type ErrorExtensions struct {
	Code              ErrorCode `json:"code"`
	WaitMilliSecondes uint      `json:"waitMilliseconds,omitempty"`
}

type Errors []Error

type response struct {
	Errors Errors          `json:"errors,omitempty"`
	Data   json.RawMessage `json:"data,omitempty"`
}

func listNoteQuery() string {
	return `query {
		notes(first: 10) {
			nodes {
				title
			}
		}				
	}`
}

// func listNoteQuery(num int, folderID ID, hasLimit bool) string {
// 	return fmt.Sprintf(`{
// 		notes(%s) {
// 			nodes {
// 				id
// 				updatedAt
// 			}
// 		}
// 	}`, buildNotesArg(num, folderID, "", hasLimit))
// }
