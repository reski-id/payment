package usecase

import (
	"errors"
	"portal/domain"
)

type topupUseCase struct {
	topupData domain.TopupData
}

func New(model domain.TopupData) domain.TopupUseCase {
	return &topupUseCase{
		topupData: model,
	}
}

func (nu *topupUseCase) AddTopup(IDUser int, newTopup domain.Topup) (domain.Topup, error) {
	if IDUser == -1 {
		return domain.Topup{}, errors.New("invalid user")
	}

	newTopup.UserID = IDUser
	res := nu.topupData.Insert(newTopup)

	if res.ID == 0 {
		return domain.Topup{}, errors.New("error insert data")
	}
	return res, nil
}
