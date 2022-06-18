package repository

import (
	"context"

	"github.com/danisbagus/simple-go-ddd/domain/entity"
)

type ProductRepository interface {
	// Insert
	Insert(ctx context.Context, product *entity.Product) error

	// Find all
	FindAll() ([]entity.Product, error)

	// Find one
	FindOneByID(ID string) (*entity.Product, error)

	// Update
	Update(ctx context.Context, ID string, product *entity.Product) error

	// Delete
	Delete(ctx context.Context, ID string) error
}
