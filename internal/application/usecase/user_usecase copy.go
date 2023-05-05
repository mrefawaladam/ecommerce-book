package usecase

import (
	"ebook/internal/adapters/repository"
	"ebook/internal/entity"
)

type BookUsecase struct {
	Repo repository.BookRepository
}

func (usecase BookUsecase) GetAllBooks() ([]entity.Book, error) {
	books, err := usecase.Repo.GetAllBooks()
	return books, err
}

func (usecase BookUsecase) GetBook(id int) (entity.Book, error) {
	user, err := usecase.Repo.GetBook(id)
	return user, err
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
