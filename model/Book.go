package model

type Book struct {
	ID       string `json:"id"`
	BookName string `json:"bookName"`
	Author   string `json:"author"`
}
