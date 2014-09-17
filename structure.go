package main

type Page struct {
	Id         int    `json:"id"`
	Test1      string `json:"test1"`
	Test2      string `json:"test2"`
	Test3      string `json:"test3"`
	Created_at string `json:"created_at"`
}

func NewPage() *Page {
	return &Page{}
}
