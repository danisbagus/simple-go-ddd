package application

import (
	"context"
	"time"

	"github.com/danisbagus/simple-go-ddd/domain/entity"
	"github.com/danisbagus/simple-go-ddd/domain/repository"
)

type ProductAppInterface interface {
	// Insert
	Insert(product *entity.Product) error

	// List
	List() ([]entity.Product, error)

	// View
	View(ID string) (*entity.Product, error)

	// Update
	Update(ID string, product *entity.Product) error

	// Delete
	Delete(ID string) error
}

type productApp struct {
	repo repository.ProductRepository
}

func NewProductApp(repo repository.ProductRepository) ProductAppInterface {
	return &productApp{repo}
}

func (s productApp) Insert(form *entity.Product) error {
	timeNow := time.Now()
	form.CreatedAt = timeNow
	err := s.repo.Insert(context.Background(), form)
	if err != nil {
		return err
	}
	return nil

}

func (s productApp) List() ([]entity.Product, error) {
	products, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s productApp) View(ID string) (*entity.Product, error) {
	product, err := s.repo.FindOneByID(ID)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s productApp) Update(ID string, form *entity.Product) error {
	_, err := s.repo.FindOneByID(ID)
	if err != nil {
		return err
	}

	err = s.repo.Update(context.Background(), ID, form)
	if err != nil {
		return err
	}
	return nil
}

func (s productApp) Delete(ID string) error {
	_, err := s.repo.FindOneByID(ID)
	if err != nil {
		return err
	}

	err = s.repo.Delete(context.Background(), ID)
	if err != nil {
		return err
	}
	return nil
}
