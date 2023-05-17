package data

import (
	"fmt"
	"portal/domain"

	"gorm.io/gorm"
)

type transferData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.TransferData {
	return &transferData{
		db: db,
	}
}

func (nd *transferData) Insert(newData domain.Transfer) domain.Transfer {
	cnv := ToLocal(newData)
	err := nd.db.Create(&cnv)
	fmt.Println("error", err.Error)
	if err.Error != nil {
		return domain.Transfer{}
	}
	return cnv.ToDomain()
}
