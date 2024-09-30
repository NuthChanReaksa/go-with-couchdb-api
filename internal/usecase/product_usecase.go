package usecase

import (
	"context"
	"errors"
	"github.com/NuthChanReaksa/go-with-couchdb-api/internal/entities"
	"github.com/NuthChanReaksa/go-with-couchdb-api/interfaces"
)

type ProductUsecase struct {
	ProductRepo interfaces.ProductRepository
}

func NewProductUsecase(repo interfaces.ProductRepository) *ProductUsecase {
	return &ProductUsecase{ProductRepo: repo}
}

func (u *ProductUsecase) Create(ctx context.Context, product *entities.Product) error {
	if product.Name == "" {
		return errors.New("name cannot be empty")
	}
	return u.ProductRepo.Create(ctx, product)
}

func (u *ProductUsecase) GetByID(ctx context.Context, id string) (*entities.Product, error) {
	return u.ProductRepo.GetByID(ctx, id)
}

func (u *ProductUsecase) Update(ctx context.Context, product *entities.Product) error {
	return u.ProductRepo.Update(ctx, product)
}

func (u *ProductUsecase) Delete(ctx context.Context, id string) error {
	return u.ProductRepo.Delete(ctx, id)
}

func (u *ProductUsecase) BulkCreateOrUpdate(ctx context.Context, products []*entities.Product) ([]entities.BulkOperationResult, error) {
	return u.ProductRepo.BulkCreateOrUpdate(ctx, products)
}
