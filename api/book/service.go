package book

import (
	"go-gin/database/models"
)

type bookService struct {
	repository BookRepository
}

func (s *bookService) GetBook() *[]models.Book {
	result := s.repository.GetBook()
	return result
}

func (s *bookService) CreateBook(book models.Book) (models.Book, error) {
	result, err := s.repository.CreateBook(book)
	return result, err
}

func (s *bookService) RetrieveBook(id string) (models.Book, error) {
	result, err := s.repository.FindById(id)
	return result, err
}

func (s *bookService) UpdateBook(id string, book models.Book) (models.Book, error) {
	result, err := s.repository.UpdateBook(id, book)
	return result, err
}

func (s *bookService) DeleteBook(id string) (models.Book, error) {
	result, err := s.repository.DeleteBook(id)
	return result, err
}
