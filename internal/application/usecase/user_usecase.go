package usecase

import (
	"ebook/internal/adapters/repository"
	"ebook/internal/entity"
)

type UserUsecase struct {
	Repo repository.UserRepository
}

func (usecase UserUsecase) GetAllUsers() ([]entity.User, error) {
	users, err := usecase.Repo.GetAllUsers()
	return users, err
}

func (usecase UserUsecase) GetUser(id int) (entity.User, error) {
	user, err := usecase.Repo.GetUser(id)
	return user, err
}

func (usecase UserUsecase) CreateUser(user entity.User) error {
	err := usecase.Repo.CreateUser(user)
	return err
}

func (usecase UserUsecase) UpdateUser(id int, user entity.User) error {
	err := usecase.Repo.UpdateUser(id, user)
	return err
}

func (usecase UserUsecase) DeleteUser(id int) error {
	err := usecase.Repo.DeleteUser(id)
	return err
}

func (usecase UserUsecase) FindUser(id int) error {
	err := usecase.Repo.FindUser(id)
	return err
}

func (usecase UserUsecase) UniqueEmail(email string) error {
	err := usecase.Repo.UniqueEmail(email)
	return err
}
func (usecase UserUsecase) GetUserByEmail(email string) (*entity.User, error) {
	return usecase.Repo.FindByEmail(email)
}
