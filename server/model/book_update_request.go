package model

type UpdateBookRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
}
