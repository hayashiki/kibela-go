package kibela

import (
	"encoding/json"
	"fmt"
)

type NoteService struct {
	client *Client
}

type Note struct {
	ID    string `json"id"`
	Title string `json"title"`
	URL   string `json"url"`
	Content string `json"content"`
}

type SearchResult struct {
	Title    string `json:"title"`
	Document Note   `json:"document"`
}

func (s *NoteService) GetAll() ([]*Note, error) {

	pa := Payload{
		Query: listNoteQuery(),
	}

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
	_, err = s.client.Do(req, &res)

	if err := json.Unmarshal(res.Data, &result); err != nil {
		return nil, err
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
        document {
          ... on Note {
            id
            title
						url
          }
				}				
			}
		}				
	}`, query)

	payload := Payload{
		Query: searchQuery,
	}

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
	_, err = s.client.Do(req, &res)

	if err := json.Unmarshal(res.Data, &result); err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return result.Search.Nodes, err
}

func (s *NoteService) Get(id string) (*Note, error) {
	const query = `
query ($id: ID!) {
  note(id: $id) {
    id
    title
    url
    content
  }
}
`
	type response struct {
		Note *Note `json:"note"`
	}
	variables := map[string]interface{}{
		"id": id,
	}
	var resp response
	err := s.client.GraphQL(query, variables, &resp)
	return resp.Note, err
}
