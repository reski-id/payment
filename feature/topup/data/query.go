package data

import (
	"portal/domain"

	"gorm.io/gorm"
)

type topupData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.TopupData {
	return &topupData{
		db: db,
	}
}

func (nd *topupData) Insert(newData domain.Topup) domain.Topup {
	cnv := ToLocal(newData)

	// Get user's current balance
	var user domain.User
	err := nd.db.First(&user, "id = ?", newData.UserID).Error
	if err != nil {
		return domain.Topup{}
	}

	// Update user's balance with new top-up amount
	user.Balance += newData.Amount
	err = nd.db.Save(&user).Error
	if err != nil {
		return domain.Topup{}
	}

	// Set balance before and balance fields in top-up data
	cnv.BalanceBefore = user.Balance - newData.Amount
	cnv.Balance = user.Balance

	err = nd.db.Create(&cnv).Error
	if err != nil {
		return domain.Topup{}
	}
	return cnv.ToDomain()
}
