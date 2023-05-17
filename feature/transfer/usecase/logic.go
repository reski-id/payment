package usecase

import (
	"errors"
	"portal/domain"
)

type transferUseCase struct {
	transferData domain.TransferData
}

func New(model domain.TransferData) domain.TransferUseCase {
	return &transferUseCase{
		transferData: model,
	}
}

func (nu *transferUseCase) AddTransfer(IDUser int, newTransfer domain.Transfer) (domain.Transfer, error) {
	if IDUser == -1 {
		return domain.Transfer{}, errors.New("invalid user")
	}

	newTransfer.UserID = IDUser
	res := nu.transferData.Insert(newTransfer)

	if res.ID == 0 {
		return domain.Transfer{}, errors.New("error insert data")
	}
	return res, nil
}
