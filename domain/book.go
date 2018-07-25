package domain

//struct Book
type Book struct {
	Model
	CategoryID  UUID   "json:categoryId"
	Name        string `json:"name"`
	Author      string `json:"author"`
	Description string `json:"description"`
}
