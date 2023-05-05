package repository

import (
	"ebook/internal/entity"

	"gorm.io/gorm"
)

type StoreRepository struct {
	DB *gorm.DB
}

func (repo StoreRepository) GetAllStores() ([]entity.Store, error) {
	var stores []entity.Store
	result := repo.DB.Find(&stores)

	return stores, result.Error
}

func (repo StoreRepository) GetStore(id int) (entity.Store, error) {
	var stores entity.Store
	result := repo.DB.First(&stores, id)

	return stores, result.Error
}

func (repo StoreRepository) CreateStore(store entity.Store) error {
	result := repo.DB.Create(&store)
	return result.Error
}

func (repo StoreRepository) UpdateStore(id int, store entity.Store) error {
	result := repo.DB.Model(&store).Where("id = ?", id).Updates(&store)
	return result.Error
}

func (repo StoreRepository) DeleteStore(id int) error {
	result := repo.DB.Delete(&entity.Store{}, id)
	return result.Error
}

func (repo StoreRepository) FindStore(id int) error {
	result := repo.DB.First(&entity.Store{}, id)
	return result.Error
}
