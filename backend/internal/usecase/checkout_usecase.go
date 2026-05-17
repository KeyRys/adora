package usecase

import (
	"backend/internal/repository"
	"fmt"
)

type CheckoutUsecase struct {
	CheckoutRepo *repository.CheckoutRepository
}

func NewCheckoutUsecase(checkoutRepo *repository.CheckoutRepository) *CheckoutUsecase {

	return &CheckoutUsecase{
		CheckoutRepo: checkoutRepo,
	}
}

func (u *CheckoutUsecase) Checkout(userID string) error {

	items, err := u.CheckoutRepo.GetCartItemsByUserID(userID)

	if err != nil {
		return err
	}

	total := 0

	for _, item := range items {
		total += item.Price
	}
	fmt.Println("totalPrice", total)

	orderID, err :=
		u.CheckoutRepo.CreateOrder(
			userID,
			total,
		)
	fmt.Println("orderID", orderID)

	if err != nil {
		return err
	}

	for _, item := range items {

		err := u.CheckoutRepo.CreateOrderItem(
			orderID,
			item,
		)

		if err != nil {
			return err
		}
	}
	fmt.Println("Created order items for user", userID)

	err = u.CheckoutRepo.ClearCart(userID)

	if err != nil {
		return err
	}

	return nil
}
