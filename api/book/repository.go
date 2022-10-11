package book

import (
	"go-gin/database"
	"go-gin/database/models"
)

type BookRepository struct {
}

func (bookRepository BookRepository) GetBook() *[]models.Book {
	var books []models.Book
	database.DB.Find(&books)
	return &books
}

func (bookRepository BookRepository) CreateBook(book models.Book) (models.Book, error) {
	err := database.DB.Create(&book).Error
	return book, err
}

func (bookRepository BookRepository) FindById(id string) (book models.Book, err error) {
	err = database.DB.Set("gorm:auto_preload", true).Where("id = ?", id).First(&book).Error
	return book, err
}

func (bookRepository BookRepository) UpdateBook(id string, book models.Book) (models.Book, error) {
	instance, err := bookRepository.FindById(id)
	if err == nil {
		instance.Description = book.Description
		instance.Name = book.Name
		database.DB.Save(&instance)
	}
	return instance, err
}
func (bookRepository BookRepository) DeleteBook(id string) (models.Book, error) {
	instance, err := bookRepository.FindById(id)
	if err == nil {
		database.DB.Delete(&instance)
	}
	return instance, err
}
