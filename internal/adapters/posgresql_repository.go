package adapters

import (
	"context"
	"crud-go/internal/models"
	"database/sql"
)

type PostgreSQLRepository struct {
	db *sql.DB
}

func NewPostgreSQLRepository(db *sql.DB) *PostgreSQLRepository {
	if db == nil {
		panic("missing db")
	}
	return &PostgreSQLRepository{db: db}
}

func (r *PostgreSQLRepository) GetAllItems(ctx context.Context) ([]models.Item, error) {
	rows, e := r.db.QueryContext(ctx, "SELECT * FROM items")
	if e != nil {
		return nil, e
	}
	defer rows.Close()

	var all []models.Item
	for rows.Next() {
		var item models.Item
		if e := rows.Scan(&item.Id, &item.Name); e != nil {
			return nil, e
		}
		all = append(all, item)
	}
	return all, nil
}
