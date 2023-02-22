package forms

type CreateBoard struct {
	Title   string `json:"title"`
	Author  string `json:"email"`
	Content string `json:"content"`
}

type BoardQuery struct {
	Id    string `form:"id"`
	Limit int    `form:"limit"`
}
type DeleteBoard struct {
	ID uint `json:"id"`
}
