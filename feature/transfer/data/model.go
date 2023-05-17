package data

import (
	"portal/domain"

	"gorm.io/gorm"
)

type Transfer struct {
	gorm.Model
	Amount        int    `json:"amount" form:"amount"`
	Remarks       string `json:"remark form:"remark`
	BalanceBefore int    `json:"balancebefore form:"balancebefore`
	Balance       int    `json:"balance form:"balance`
	UserID        int
	User          User `gorm:"foreignKey:UserID; references:ID; constraint:OnDelete:CASCADE"`
}

type User struct {
	gorm.Model
}

func (b *Transfer) ToDomain() domain.Transfer {
	return domain.Transfer{
		ID:            int(b.ID),
		Amount:        b.Amount,
		Remarks:       b.Remarks,
		Balance:       b.Balance,
		BalanceBefore: b.BalanceBefore,
		CreatedAt:     b.CreatedAt,
		UpdatedAt:     b.UpdatedAt,
		UserID:        b.UserID,
	}
}

func ParseToArr(arr []Transfer) []domain.Transfer {
	var res []domain.Transfer

	for _, val := range arr {
		res = append(res, val.ToDomain())
	}
	return res
}

func ToLocal(data domain.Transfer) Transfer {
	var res Transfer
	res.ID = uint(data.ID)
	res.UserID = data.UserID
	res.Amount = data.Amount
	res.Remarks = data.Remarks
	res.Balance = data.Balance
	res.BalanceBefore = data.BalanceBefore
	return res
}
