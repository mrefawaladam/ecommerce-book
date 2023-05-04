package repository

import (
	"ebook/internal/entity"
	"fmt"

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

func (repo BookRepository) GetUser(id int) (entity.User, error) {
	var users entity.User
	result := repo.DB.Preload("Blogs", "deleted_at IS NULL").First(&users, id)
	// result := repo.DB.First(&users, id)

	return users, result.Error
}

func (repo BookRepository) CreateUser(user entity.User) error {
	result := repo.DB.Create(&user)
	return result.Error
}

func (repo BookRepository) UpdateUser(id int, user entity.User) error {
	result := repo.DB.Model(&user).Where("id = ?", id).Updates(&user)
	return result.Error
}

func (repo BookRepository) DeleteUser(id int) error {
	result := repo.DB.Delete(&entity.User{}, id)
	return result.Error
}

func (repo BookRepository) FindUser(id int) error {
	result := repo.DB.First(&entity.User{}, id)
	return result.Error
}
func (repo BookRepository) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := repo.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (repo BookRepository) UniqueEmail(email string) error {
	var user entity.User
	result := repo.DB.Where("email = ?", email).First(&user)
	if result.RowsAffected > 0 {
		return fmt.Errorf("email %s already exists", email)
	}
	return nil
}
