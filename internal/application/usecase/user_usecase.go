package usecase

import (
	"ebook/internal/adapters/repository"
	"ebook/internal/entity"
)

type Usecase struct {
	Repo repository.Repository
}

func (usecase Usecase) GetAllUsers() ([]entity.User, error) {
	users, err := usecase.Repo.GetAllUsers()
	return users, err
}

func (usecase Usecase) GetUser(id int) (entity.User, error) {
	user, err := usecase.Repo.GetUser(id)
	return user, err
}

func (usecase Usecase) CreateUser(user entity.User) error {
	err := usecase.Repo.CreateUser(user)
	return err
}

func (usecase Usecase) UpdateUser(id int, user entity.User) error {
	err := usecase.Repo.UpdateUser(id, user)
	return err
}

func (usecase Usecase) DeleteUser(id int) error {
	err := usecase.Repo.DeleteUser(id)
	return err
}

func (usecase Usecase) SearchUser(id int) error {
	err := usecase.Repo.SearchUser(id)
	return err
}

func (usecase Usecase) UniqueEmail(email string) error {
	err := usecase.Repo.UniqueEmail(email)
	return err
}
