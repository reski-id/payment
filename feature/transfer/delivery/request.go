package delivery

import (
	"portal/domain"
)

type InsertRequest struct {
	Amount        int    `json:"amount" form:"amount"`
	Remarks       string `json:"remark form:"remark`
	BalanceBefore int    `json:"balancebefore form:"balancebefore`
	Balance       int    `json:"balance form:"balance`
	UserID        int    `json:"user_id" form:"user_id"`
}

func (ni *InsertRequest) ToDomain() domain.Transfer {
	return domain.Transfer{
		Amount:        ni.Amount,
		BalanceBefore: ni.BalanceBefore,
		Balance:       ni.Balance,
		UserID:        ni.UserID,
	}
}
