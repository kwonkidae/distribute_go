package booksHandler

import "fmt"

type bookResource struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Link  string `json:"link"`
}

type requestPayload struct {
	Title string `json:"title"`
	Link  string `json:"link"`
}

type response struct {
	StatusCode int
	Books      []bookResource
}

type Action struct {
	Id      string
	Type    string
	Payload requestPayload
	RetChan chan<- response
}

func GetBooks() map[string]bookResource {
	books := map[string]bookResource{}
	for i := 1; i < 6; i++ {
		id := fmt.Sprintf("%d", i)
		books[id] = bookResource{
			Id:    id,
			Title: fmt.Sprintf("Book-%s", id),
			Link:  fmt.Sprintf("http://link-to-book%s.com", id),
		}
	}
	return books
}
