package repository

import (
	"backend/internal/domain"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type SellerRepository struct {
	DB *pgx.Conn
}

func NewSellerRepository(db *pgx.Conn) *SellerRepository {
	return &SellerRepository{
		DB: db,
	}
}

func (r *SellerRepository) IsSeller(userID string) (bool, error) {
	var exists bool

	query := `
		SELECT EXISTS(
			SELECT 1
			FROM sellers
			WHERE user_id = $1::uuid
		)
	`

	err := r.DB.QueryRow(context.Background(), query, userID).Scan(&exists)

	return exists, err
}

func (r *SellerRepository) CreateSeller(userID, location string) error {
	query := `
		INSERT INTO sellers (user_id, location)
		VALUES ($1::uuid, $2)
	`

	_, err := r.DB.Exec(
		context.Background(),
		query,
		userID,
		location,
	)

	return err
}

func (r *SellerRepository) GetSellerIDByUserID(userID string) (string, error) {

	var sellerID string

	query := `
		SELECT id
		FROM sellers
		WHERE user_id = $1::uuid
	`

	err := r.DB.QueryRow(
		context.Background(),
		query,
		userID,
	).Scan(&sellerID)

	return sellerID, err
}

func (r *SellerRepository) CreateRabbit(rabbit *domain.Rabbit) error {
	rabbit.HealthStatus = "healthy"
	query := `
		INSERT INTO rabbits (
			seller_id,
			name,
			breed,
			gender,
			age,
			weight,
			color,
			price,
			purpose,
			description,
			health_status
		)
		VALUES (
			$1::uuid,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7,
			$8,
			$9,
			$10,
			$11
		)
	`
	fmt.Printf("Creating rabbit: %+v\n", rabbit)
	_, err := r.DB.Exec(
		context.Background(),
		query,
		rabbit.SellerID,
		rabbit.Name,
		rabbit.Breed,
		rabbit.Gender,
		rabbit.Age,
		rabbit.Weight,
		rabbit.Color,
		rabbit.Price,
		rabbit.Purpose,
		rabbit.Description,
		rabbit.HealthStatus,
		//rabbit.ImageURL,
	)
	fmt.Printf("Error creating rabbit: %v\n", err)
	return err
}

func (r *SellerRepository) GetSellerRabbits(userID string) ([]domain.Rabbit, error) {
	//fmt.Println("Getting rabbits for user ID:", userID)
	ctx := context.Background()
	query := `
		SELECT
			r.id,
			r.seller_id,
			r.name,
			r.breed,
			r.age,
			r.gender,
			r.weight,
			r.color,
			r.price,
			r.purpose,
			r.health_status,
			r.description
		FROM rabbits r
		JOIN sellers s
			ON r.seller_id = s.id
		WHERE s.user_id = $1::uuid
	`

	rows, err := r.DB.Query(
		ctx,
		query,
		userID,
	)

	if err != nil {
		fmt.Println("Error querying rabbits:", err)
		return nil, err
	}

	defer rows.Close()

	var rabbits []domain.Rabbit

	for rows.Next() {

		var rabbit domain.Rabbit

		err := rows.Scan(
			&rabbit.ID,
			&rabbit.SellerID,
			&rabbit.Name,
			&rabbit.Breed,
			&rabbit.Age,
			&rabbit.Gender,
			&rabbit.Weight,
			&rabbit.Color,
			&rabbit.Price,
			&rabbit.Purpose,
			&rabbit.HealthStatus,
			&rabbit.Description,
			//&rabbit.ImageURL,
			//&rabbit.Status,
		)

		if err != nil {
			return nil, err
		}

		rabbits = append(rabbits, rabbit)
	}

	return rabbits, nil
}

func (r *SellerRepository) UpdateRabbit(
	rabbitID string,
	sellerID string,
	name string,
	breed string,
	healthstatus string,
	price int,
	description string,
) error {
	fmt.Println("Getting rabbits for user ID:", sellerID)
	ctx := context.Background()

	_, err := r.DB.Exec(ctx, `
        UPDATE rabbits
        SET
            name = $1,
            breed = $2,
            price = $3,
            description = $4,
			health_status = $5
        WHERE id = $6
        AND seller_id = $7
    `,
		name,
		breed,
		price,
		description,
		healthstatus,
		rabbitID,
		sellerID,
	)

	return err
}

func (r *SellerRepository) DeleteRabbit(
	rabbitID string,
	sellerID string,
) error {

	ctx := context.Background()

	_, err := r.DB.Exec(ctx, `
        DELETE FROM rabbits
        WHERE id = $1
        AND seller_id = $2
    `,
		rabbitID,
		sellerID,
	)

	return err
}
