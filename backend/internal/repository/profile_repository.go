package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type ProfileRepository struct {
	DB *pgx.Conn
}

func NewProfileRepository(db *pgx.Conn) *ProfileRepository {
	return &ProfileRepository{
		DB: db,
	}
}

type ProfileResponse struct {
	ID       string `json:"id"`
	UserID   string `json:"user_id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Isseller bool   `json:"is_seller"`
}

func (r *ProfileRepository) GetProfileByUserID(
	userID string,
) (*ProfileResponse, error) {

	var profile ProfileResponse

	err := r.DB.QueryRow(context.Background(), `
		SELECT
			id,
			user_id,
			name,
			phone,
			address,
			EXISTS (
				SELECT *
				FROM sellers
				WHERE sellers.user_id = profiles.user_id
			) AS is_seller
		FROM profiles
		WHERE user_id = $1::uuid
	`,
		userID,
	).Scan(
		&profile.ID,
		&profile.UserID,
		&profile.Name,
		&profile.Phone,
		&profile.Address,
		&profile.Isseller,
	)

	if err != nil {
		return nil, err
	}

	return &profile, nil
}
