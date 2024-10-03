package repository

import (
	"context"
	"github.com/NuthChanReaksa/go-with-couchdb-api/internal/entities"
	"github.com/NuthChanReaksa/go-with-couchdb-api/internal/interfaces"
	"github.com/go-kivik/kivik/v4"
	_ "github.com/go-kivik/kivik/v4/couchdb"
)

type CouchDBProductRepository struct {
	db *kivik.DB
}

func NewCouchDBProductRepository(db *kivik.DB) interfaces.ProductRepository {
	return &CouchDBProductRepository{db: db}
}

func (r *CouchDBProductRepository) Create(ctx context.Context, product *entities.Product) error {
	_, err := r.db.Put(ctx, product.ID, product)
	return err
}

func (r *CouchDBProductRepository) GetByID(ctx context.Context, id string) (*entities.Product, error) {
	product := new(entities.Product)
	err := r.db.Get(ctx, id).ScanDoc(product)
	return product, err
}

func (r *CouchDBProductRepository) Update(ctx context.Context, product *entities.Product) error {
	rev, err := r.db.GetRev(ctx, product.ID)
	if err != nil {
		return err
	}
	_, err = r.db.Put(ctx, product.ID, product, kivik.Rev(rev))
	return err
}

func (r *CouchDBProductRepository) Delete(ctx context.Context, id string) error {
	rev, err := r.db.GetRev(ctx, id)
	if err != nil {
		return err
	}
	_, err = r.db.Delete(ctx, id, rev)
	return err
}

func (r *CouchDBProductRepository) BulkCreateOrUpdate(ctx context.Context, products []*entities.Product) ([]entities.BulkOperationResult, error) {
	bulkDocs := make([]interface{}, len(products))
	for i, product := range products {
		bulkDocs[i] = product
	}

	results, err := r.db.BulkDocs(ctx, bulkDocs)
	if err != nil {
		return nil, err
	}

	operationResults := make([]entities.BulkOperationResult, len(results))
	for i, result := range results {
		operationResults[i] = entities.BulkOperationResult{
			ID:    result.ID,
			Rev:   result.Rev,
			Error: result.Error,
		}
	}
	return operationResults, nil
}
