package interfaces

import (
	"context"
	"cmd/internal/entities"
)

type ProductRepository interface {
	Create(ctx context.Context, product *entities.Product) error
	GetByID(ctx context.Context, id string) (*entities.Product, error)
	Update(ctx context.Context, product *entities.Product) error
	Delete(ctx context.Context, id string) error
	BulkCreateOrUpdate(ctx context.Context, products []*entities.Product) ([]entities.BulkOperationResult, error)
}
