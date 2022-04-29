package kibela

import "log"

type AttachmentService struct {
	client *Client
}

type Attachment struct {
	ID string `json:"id"`
	Key string `json:"key"`
	Kind string `json:"kind"`
	Name string `json:"name"`
	Data []byte `json:"data"`
	DataUrl string `json:"dataUrl"`
	MimeType string `json:"mimeType"`
}

func (s *AttachmentService) Get(path string) (*Attachment, error){
	const query = `
query ($path: String!) {
  attachmentFromPath(path: $path) {
    id
    key
    kind
    name
    mimeType
    dataUrl
  }
}
`
	type response struct {
		Attachment *Attachment `json:"attachmentFromPath"`
	}

	variables := map[string]interface{}{
		"path": path,
	}

	var resp response

	err := s.client.GraphQL(query, variables, &resp)
	log.Printf("resp is %v", resp.Attachment.ID)
	log.Printf("err %v", err)
	return resp.Attachment, nil
}
