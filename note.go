package kibela

import (
	"encoding/json"
	"fmt"
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

type SearchResult struct {
	// ID `json:"id"`
	Title string `json:"title"`
	// Content string `json"content"`
}

func (s *NoteService) GetAll() ([]*Note, error) {

	pa := Payload{
		Query: listNoteQuery(),
	}

	pa.Query = strings.TrimSpace(pa.Query)

	log.Printf("log pa %v", pa)

	req, err := s.client.NewRequest("POST", pa)
	if err != nil {
		return nil, err
	}

	var result struct {
		Notes struct {
			Nodes []*Note `json:"nodes"`
		} `json:"notes"`
	}

	var res Response

	// noteResp := interface{}
	resp, err := s.client.Do(req, &res)

	log.Printf("log res %v", resp)

	if err := json.Unmarshal(res.Data, &result); err != nil {
		return nil, err
	}

	for _, note := range result.Notes.Nodes {
		log.Printf("log res %v", note.Title)
	}

	if err != nil {
		return nil, err
	}

	return result.Notes.Nodes, err
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

func (s *NoteService) Search(query string) ([]*SearchResult, error) {

	searchQuery := fmt.Sprintf(`query {
		search(first: 10, query: "%s") {
			nodes {
				title
			}
		}				
	}`, query)

	payload := Payload{
		Query: searchQuery,
	}

	// pa.Query = strings.TrimSpace(pa.Query)

	req, err := s.client.NewRequest("POST", payload)
	if err != nil {
		return nil, err
	}

	var result struct {
		Search struct {
			Nodes []*SearchResult `json:"nodes"`
		} `json:"search"`
	}

	var res Response
	resp, err := s.client.Do(req, &res)

	log.Printf("log resp %v", resp)

	if err := json.Unmarshal(res.Data, &result); err != nil {
		return nil, err
	}

	for _, note := range result.Search.Nodes {
		log.Printf("log res %v", note.Title)
	}

	if err != nil {
		return nil, err
	}

	return result.Search.Nodes, err
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
