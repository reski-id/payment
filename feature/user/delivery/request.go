package delivery

import (
	"portal/domain"
)

type InsertFormat struct {
	FirstName   string `json:"first_name" form:"first_name" validate:"required"`
	LastName    string `json:"last_name" form:"last_name" validate:"required"`
	PhoneNumber string `json:"phone_number" form:"phone_number" validate:"required"`
	Address     string `json:"address" form:"address" validate:"required"`
	Pin         string `json:"pin" form:"pin"`
}

func (u *InsertFormat) ToModel() domain.User {
	return domain.User{
		FirstName:   u.FirstName,
		LastName:    u.LastName,
		PhoneNumber: u.PhoneNumber,
		Address:     u.Address,
	}
}

type LoginFormat struct {
	PhoneNumber string `json:"phone_number" form:"phone_number"`
	Pin         string `json:"pin" form:"pin"`
}

func (lf *LoginFormat) LoginToModel() domain.User {
	return domain.User{
		PhoneNumber: lf.PhoneNumber,
		Pin:         lf.Pin,
	}
}
