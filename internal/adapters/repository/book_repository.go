package repository

import (
	"ebook/internal/entity"

	"gorm.io/gorm"
)

type BookRepository struct {
	DB *gorm.DB
}

func (repo BookRepository) GetAllBooks() ([]entity.Book, error) {
	var books []entity.Book
	result := repo.DB.Preload("Category", "deleted_at IS NULL").Find(&books)

	return books, result.Error
}

func (repo BookRepository) GetBook(id int) (entity.Book, error) {
	var books entity.Book
	result := repo.DB.Preload("Category", "deleted_at IS NULL").First(&books, id)
	return books, result.Error
}

func (repo BookRepository) CreateBook(book entity.Book) error {
	result := repo.DB.Create(&book)
	return result.Error
}

func (repo BookRepository) UpdateBook(id int, book entity.Book) error {
	result := repo.DB.Model(&book).Where("id = ?", id).Updates(&book)
	return result.Error
}

func (repo BookRepository) DeleteBook(id int) error {
	result := repo.DB.Delete(&entity.Book{}, id)
	return result.Error
}

func (repo BookRepository) FindBook(id int) error {
	result := repo.DB.First(&entity.Book{}, id)
	return result.Error
}
