package forms

type CreateBoard struct {
	Title   string `json:"title"`
	Author  string `json:"email"`
	Content string `json:"content"`
}
