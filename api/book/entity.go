package book

type RequestPostBook struct {
	Description string `json:"description" binding:"required"`
	Name        string `json:"name" binding:"required"`
}
