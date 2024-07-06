package usecase

import (
	"checkout/model"
	"checkout/repository"
)

type CheckoutUsecase struct {
	repository repository.CheckoutRepository
}

func NewCheckoutUsecase(repository repository.CheckoutRepository) CheckoutUsecase {
	return CheckoutUsecase{
		repository: repository,
	}
}

func (cu *CheckoutUsecase) Checkout(order model.Order) ([]byte, error) {
	return cu.repository.Checkout(order)
}
