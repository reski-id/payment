package delivery

import "portal/domain"

type DataResponse struct {
	ID            int    `json:"id"`
	Amount        int    `json:"amount"`
	Remarks       string `json:"remark"`
	BalanceBefore int    `json:"balancebefore"`
	Balance       int    `json:"balance"`
	UserID        int    `json:"user_id"`
}

func FromDomain(data domain.Payment) DataResponse {
	var res DataResponse
	res.ID = int(data.ID)
	res.Amount = data.Amount
	res.Remarks = data.Remarks
	res.Balance = data.Balance
	res.BalanceBefore = data.BalanceBefore
	res.UserID = int(data.UserID)
	return res
}
