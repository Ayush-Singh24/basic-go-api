package product

import (
	"database/sql"
	"fmt"

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
		&product.Description,
		&product.Image,
		&product.Price,
		&product.Quantity,
		&product.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *Store) CreateProduct(p types.Product) error {
	_, err := s.db.Exec("INSERT INTO products (name, description, image, price, quantity) VALUES (?,?,?,?,?)", p.Name, p.Description, p.Image, p.Price, p.Quantity)

	if err != nil {
		return err
	}
	return nil
}

func (s *Store) GetProductById(id int) (*types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products WHERE id=?", id)

	if err != nil {
		return nil, err
	}

	product := new(types.Product)

	for rows.Next() {
		product, err = scanRowIntoProduct(rows)
		if err != nil {
			return nil, err
		}

	}

	if product.Id == 0 {
		return nil, fmt.Errorf("product not found")
	}

	return product, nil
}
