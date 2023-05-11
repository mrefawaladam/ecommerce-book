package usecase

import (
	"ebook/internal/adapters/repository"
	"ebook/internal/entity"
	"strings"
)

type BookUsecase struct {
	Repo repository.BookRepository
}

func (usecase BookUsecase) GetAllBooks() ([]entity.Book, error) {
	books, err := usecase.Repo.GetAllBooks()
	return books, err
}

func (usecase BookUsecase) GetBook(id int) (entity.Book, error) {
	book, err := usecase.Repo.GetBook(id)
	return book, err
}

func (usecase BookUsecase) CreateBook(user entity.Book) error {
	err := usecase.Repo.CreateBook(user)
	return err
}

func (usecase BookUsecase) UpdateBook(id int, book entity.Book) error {
	err := usecase.Repo.UpdateBook(id, book)
	return err
}

func (usecase BookUsecase) DeleteBook(id int) error {
	err := usecase.Repo.DeleteBook(id)
	return err
}

func (usecase BookUsecase) FindBook(id int) error {
	err := usecase.Repo.FindBook(id)
	return err
}
func (bu *BookUsecase) SearchBooks(searchQuery string) ([]entity.Book, error) {
	books, err := bu.Repo.GetAllBooks()
	if err != nil {
		return nil, err
	}

	var searchResults []entity.Book

	for _, book := range books {
		if strings.Contains(strings.ToLower(book.Title), strings.ToLower(searchQuery)) ||
			strings.Contains(strings.ToLower(book.Author), strings.ToLower(searchQuery)) {
			searchResults = append(searchResults, book)
		}
	}

	return searchResults, nil
}
