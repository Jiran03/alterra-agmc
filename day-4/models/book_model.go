package models

type Book struct {
	ID              int    `json:"id" form:"id"`
	Title           string `json:"title" form:"title"`
	Writer          string `json:"writer" form:"writer"`
	Publisher       string `json:"publisher" form:"publisher"`
	PublicationYear int    `json:"publication_year" form:"publication_year"`
}
