package usecase

import (
	"backend/internal/domain"
	"backend/internal/repository"
	"errors"
	"fmt"
)

type SellerUsecase struct {
	Repo *repository.SellerRepository
}

func NewSellerUsecase(r *repository.SellerRepository) *SellerUsecase {
	return &SellerUsecase{
		Repo: r,
	}
}

func (u *SellerUsecase) BecomeSeller(userID, location string) error {

	exists, err := u.Repo.IsSeller(userID)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("user already seller")
	}

	return u.Repo.CreateSeller(userID, location)
}

func (u *SellerUsecase) CreateRabbit(
	userID string,
	rabbit *domain.Rabbit,
) error {

	sellerID, err := u.Repo.GetSellerIDByUserID(userID)

	if err != nil {
		fmt.Printf("Error getting seller ID for user %s: %v\n", userID, err)
		fmt.Printf("Rabbit data: %+v\n", rabbit)
		return err
	}

	rabbit.SellerID = sellerID

	return u.Repo.CreateRabbit(rabbit)
}

func (u *SellerUsecase) GetSellerRabbits(userID string) ([]domain.Rabbit, error) {
	return u.Repo.GetSellerRabbits(userID)
}

func (u *SellerUsecase) UpdateRabbit(
	rabbitID string,
	userID string,
	name string,
	breed string,
	healthstatus string,
	price int,
	description string,
) error {

	sellerID, err := u.Repo.GetSellerIDByUserID(userID)
	if err != nil {
		return err
	}

	return u.Repo.UpdateRabbit(
		rabbitID,
		sellerID,
		name,
		breed,
		healthstatus,
		price,
		description,
	)
}

func (u *SellerUsecase) DeleteRabbit(
	rabbitID string,
	userID string,
) error {

	sellerID, err := u.Repo.GetSellerIDByUserID(userID)
	if err != nil {
		return err
	}

	return u.Repo.DeleteRabbit(
		rabbitID,
		sellerID,
	)
}
