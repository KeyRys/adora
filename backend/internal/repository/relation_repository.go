package repository

import (
	"context"
	"fmt"

	"backend/internal/domain"

	"github.com/jackc/pgx/v5"
)

type RabbitRelationRepository struct {
	DB *pgx.Conn
}

func NewRabbitRelationRepository(db *pgx.Conn) *RabbitRelationRepository {

	return &RabbitRelationRepository{
		DB: db,
	}
}

func (r *RabbitRelationRepository) CreateRelation(relation domain.RabbitRelation) error {

	_, err := r.DB.Exec(
		context.Background(),
		`
		INSERT INTO rabbit_relationships (
			parent_id,
			child_id
		)
		VALUES (
			$1::uuid,
			$2::uuid
		)
		`,
		relation.ParentID,
		relation.ChildID,
	)

	return err
}

func (r *RabbitRelationRepository) GetSellerRelations(userID string) ([]domain.RabbitRelationResponse, error) {

	rows, err := r.DB.Query(
		context.Background(),
		`
		SELECT
			parent.name parent_name,
			child.name child_name

		FROM rabbit_relationships rr

		JOIN rabbits parent
			ON rr.parent_id = parent.id

		JOIN rabbits child
			ON rr.child_id = child.id

		JOIN sellers s
			ON child.seller_id = s.id

		WHERE s.user_id = $1::uuid
		`,
		userID,
	)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer rows.Close()

	var relations []domain.RabbitRelationResponse

	for rows.Next() {

		var relation domain.RabbitRelationResponse

		err := rows.Scan(
			&relation.ParentName,
			&relation.ChildName,
		)

		if err != nil {
			return nil, err
		}

		relations = append(
			relations,
			relation,
		)
	}
	fmt.Println(relations)
	return relations, nil
}
