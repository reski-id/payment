package usecase

import (
	"errors"
	"portal/domain"
)

type paymentUseCase struct {
	paymentData domain.PaymentData
}

func New(model domain.PaymentData) domain.PaymentUseCase {
	return &paymentUseCase{
		paymentData: model,
	}
}

func (nu *paymentUseCase) AddPayment(IDUser int, newPayment domain.Payment) (domain.Payment, error) {
	if IDUser == -1 {
		return domain.Payment{}, errors.New("invalid user")
	}

	newPayment.UserID = IDUser
	res := nu.paymentData.Insert(newPayment)

	if res.ID == 0 {
		return domain.Payment{}, errors.New("error insert data")
	}
	return res, nil
}
