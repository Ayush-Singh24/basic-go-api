package product

import (
	"database/sql"

	"github.com/Ayush-Singh24/basic-go-api/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetProducts() ([]types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}

	products := make([]types.Product, 0)

	for rows.Next() {
		p, err := scanRowIntoProduct(rows)
		if err != nil {
			return nil, err
		}

		products = append(products, *p)
	}

	return products, nil
}

func scanRowIntoProduct(rows *sql.Rows) (*types.Product, error) {
	product := new(types.Product)
	err := rows.Scan(
		&product.Id,
		&product.Name,
		&product.Price,
		&product.Description,
		&product.Quantity,
		&product.Image,
		&product.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return product, nil
}
