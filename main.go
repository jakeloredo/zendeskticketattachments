package zendeskticketattachments

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

// Zendesk is the struct that holds zendesk domain and credentials
type Zendesk struct {
	Subdomain string
	Username  string
	Password  string
}

// Attachment represents an attachment object from Zendesk ticket comments json
type Attachment struct {
	FileName   string `json:"file_name"`
	ContentURL string `json:"content_url"`
}

// Comment represents a comment object from Zendesk ticket comments json
type Comment struct {
	ID          int          `json:"id"`
	Attachments []Attachment `json:"attachments"`
}

// ResponseStruct is used to unmarshal the json from the response to an API request
type ResponseStruct struct {
	Comments []Comment `json:"comments"`
}

func (zd *Zendesk) get(endpoint string) ResponseStruct {
	urlComponents := []string{"https://", zd.Subdomain, ".zendesk.com/api/v2/", endpoint}
	url := strings.Join(urlComponents, "")
	client := http.Client{Timeout: time.Second * 4}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("content-type", "application/json")
	req.SetBasicAuth(zd.Username, zd.Password)
	res, resErr := client.Do(req)

	if resErr != nil {
		log.Fatal(resErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	//fmt.Println(string(body))
	data := ResponseStruct{}
	jsonErr := json.Unmarshal(body, &data)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return data
}

// GetTicketComments returns an array of all the comments in the specified ticket number
func (zd *Zendesk) GetTicketComments(ticketNumber string) []Comment {
	endpointComponents := []string{"tickets/", ticketNumber, "/comments"}
	endpoint := strings.Join(endpointComponents, "")
	comments := zd.get(endpoint).Comments
	return comments
}

// DownloadAttachments downloads all of the attachments in the specified ticket number
func (zd *Zendesk) DownloadAttachments(ticketNumber string) {
	r := zd.GetTicketComments(ticketNumber)

	for _, comment := range r {
		for _, attachment := range comment.Attachments {
			err := DownloadFile(attachment.FileName, attachment.ContentURL)
			if err != nil {
				panic(err)
			}
		}
	}
}
