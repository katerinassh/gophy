package repository

import (
	"context"
	"crud-go/internal/models"
)

type Repository interface {
	GetAllItems(ctx context.Context) ([]models.Item, error)
	// GetItemById(ctx context.Context, id int) (models.Item, error)
	// Create(ctx context.Context, item models.Item) (*models.Item, error)
	// Update(ctx context.Context, id int, updated models.Item) (*models.Item, error)
	// Delete(ctx context.Context, id int) error
}
