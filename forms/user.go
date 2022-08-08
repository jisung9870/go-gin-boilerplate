package forms

type UserForm struct{}

type Register struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
