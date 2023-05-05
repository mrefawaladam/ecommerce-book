package usecase

import (
	"ebook/internal/adapters/repository"
	"ebook/internal/entity"
)

type StoreUsecase struct {
	Repo repository.StoreRepository
}

func (usecase StoreUsecase) GetAllStores() ([]entity.Store, error) {
	stores, err := usecase.Repo.GetAllStores()
	return stores, err
}

func (usecase StoreUsecase) GetStore(id int) (entity.Store, error) {
	store, err := usecase.Repo.GetStore(id)
	return store, err
}

func (usecase StoreUsecase) CreateStore(store entity.Store) error {
	err := usecase.Repo.CreateStore(store)
	return err
}

func (usecase StoreUsecase) UpdateStore(id int, store entity.Store) error {
	err := usecase.Repo.UpdateStore(id, store)
	return err
}

func (usecase StoreUsecase) DeleteStore(id int) error {
	err := usecase.Repo.DeleteStore(id)
	return err
}

func (usecase StoreUsecase) FindStore(id int) error {
	err := usecase.Repo.FindStore(id)
	return err
}
