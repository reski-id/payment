package delivery

import "portal/domain"

type DataResponse struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	Pin         string `json:"pin"`
}

func FromDomain(data domain.User) DataResponse {
	var res DataResponse
	res.ID = int(data.ID)
	res.FirstName = data.FirstName
	res.LastName = data.LastName
	res.PhoneNumber = data.PhoneNumber
	res.Address = data.Address
	res.Pin = data.Pin
	return res
}
