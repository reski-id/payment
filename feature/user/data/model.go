package data

import (
	"portal/domain"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName     string `json:"first_name" form:"first_name" validate:"required"`
	LastName      string `json:"last_name" form:"last_name" validate:"required"`
	PhoneNumber   string `json:"phone_number" form:"phone_number" validate:"required"`
	Address       string `json:"address" form:"address" validate:"required"`
	Pin           string `json:"pin" form:"pin"`
	BalanceBefore int    `json:"balancebefore form:"balancebefore`
	Balance       int    `json:"balance form:"balance`
}

type UserLog struct {
	gorm.Model
	BalanceBefore   int    `json:"balancebefore form:"balancebefore`
	Balance         int    `json:"balance form:"balance`
	TransactionType int    `json:"type" form:"type" validate:"required"`
	TrannsctionLog  string `json:"log" form:"log" validate:"required"`
}

func (u *User) ToDomain() domain.User {
	return domain.User{
		ID:            int(u.ID),
		FirstName:     u.FirstName,
		LastName:      u.LastName,
		PhoneNumber:   u.PhoneNumber,
		Address:       u.Address,
		Pin:           u.Pin,
		Balance:       u.Balance,
		BalanceBefore: u.BalanceBefore,
		CreatedAt:     u.CreatedAt,
		UpdatedAt:     u.UpdatedAt,
	}
}

func (u *User) ToDomain2() domain.User {
	return domain.User{
		ID:            int(u.ID),
		FirstName:     u.FirstName,
		LastName:      u.LastName,
		PhoneNumber:   u.PhoneNumber,
		Balance:       u.Balance,
		BalanceBefore: u.BalanceBefore,
		Pin:           u.Pin,
		CreatedAt:     u.CreatedAt,
		UpdatedAt:     u.UpdatedAt,
	}
}

func ParseToArr(arr []User) []domain.User {
	var res []domain.User

	for _, val := range arr {
		res = append(res, val.ToDomain())
	}

	return res
}

func ParseToArr2(arr []User) []domain.User {
	var res []domain.User

	for _, val := range arr {
		res = append(res, val.ToDomain2())
	}

	return res
}

func ToLocal(data domain.User) User {
	var res User
	res.FirstName = data.FirstName
	res.LastName = data.LastName
	res.Pin = data.Pin
	res.PhoneNumber = data.PhoneNumber
	res.Address = data.Address
	return res
}
