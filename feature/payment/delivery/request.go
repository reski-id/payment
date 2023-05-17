package delivery

import (
	"portal/domain"
)

type InsertRequest struct {
	Amount        int    `json:"amount" form:"amount"`
	Remarks       string `json:"remark" form:"remark"`
	BalanceBefore int    `json:"balancebefore" form:"balancebefore"`
	Balance       int    `json:"balance" form:"balance"`
	UserID        int    `json:"user_id" form:"user_id"`
}

func (ni *InsertRequest) ToDomain() domain.Payment {
	return domain.Payment{
		Amount:        ni.Amount,
		Remarks:       ni.Remarks,
		BalanceBefore: ni.BalanceBefore,
		Balance:       ni.Balance,
		UserID:        ni.UserID,
	}
}
