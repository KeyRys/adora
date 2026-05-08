package repository

import (
	"backend/internal/domain"
	"context"

	"github.com/jackc/pgx/v5"
)

type ProductRepository struct {
	DB *pgx.Conn
}

func NewProductRepository(db *pgx.Conn) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (r *ProductRepository) GetAll() ([]domain.Product, error) {
	rows, err := r.DB.Query(context.Background(), "SELECT id, name, breed, price FROM rabbits")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []domain.Product

	for rows.Next() {
		var p domain.Product
		err := rows.Scan(&p.ID, &p.Name, &p.Breed, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

func (r *ProductRepository) GetByID(id string) (*domain.Product, error) {
	row := r.DB.QueryRow(context.Background(), "SELECT id, seller_id, name, breed, weight, color, gender, price FROM rabbits WHERE id = $1", id)

	var p domain.Product
	err := row.Scan(&p.ID, &p.SellerID, &p.Name, &p.Breed, &p.Weight, &p.Color, &p.Gender, &p.Price)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
