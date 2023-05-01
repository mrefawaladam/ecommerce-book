package repository

import (
	"ebook/internal/entity"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (repo Repository) GetAllUsers() ([]entity.User, error) {
	var users []entity.User
	result := repo.DB.Preload("Blogs", "deleted_at IS NULL").Find(&users)
	// result := repo.DB.Find(&users)

	return users, result.Error
}

func (repo Repository) GetUser(id int) (entity.User, error) {
	var users entity.User
	result := repo.DB.Preload("Blogs", "deleted_at IS NULL").First(&users, id)
	// result := repo.DB.First(&users, id)

	return users, result.Error
}

func (repo Repository) CreateUser(user entity.User) error {
	result := repo.DB.Create(&user)
	return result.Error
}

func (repo Repository) UpdateUser(id int, user entity.User) error {
	result := repo.DB.Model(&user).Where("id = ?", id).Updates(&user)
	return result.Error
}

func (repo Repository) DeleteUser(id int) error {
	result := repo.DB.Delete(&entity.User{}, id)
	return result.Error
}

func (repo Repository) SearchUser(id int) error {
	result := repo.DB.First(&entity.User{}, id)
	return result.Error
}
