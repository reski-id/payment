package data

import (
	"fmt"
	"portal/domain"

	"gorm.io/gorm"
)

type paymentData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.PaymentData {
	return &paymentData{
		db: db,
	}
}

func (nd *paymentData) Insert(newData domain.Payment) domain.Payment {
	cnv := ToLocal(newData)
	err := nd.db.Create(&cnv)
	fmt.Println("error", err.Error)
	if err.Error != nil {
		return domain.Payment{}
	}
	return cnv.ToDomain()
}
