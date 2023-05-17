package domain

import (
	"time"

	"github.com/labstack/echo/v4"
)

type User struct {
	ID            int
	FirstName     string
	LastName      string
	PhoneNumber   string
	Address       string
	Pin           string
	Balance       int
	BalanceBefore int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type UserLog struct {
	ID              int
	Balance         int
	BalanceBefore   int
	TransactionType int
	TrannsctionLog  string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type UserHandler interface {
	InsertUser() echo.HandlerFunc
	LoginUser() echo.HandlerFunc
	GetProfile() echo.HandlerFunc
	DeleteUser() echo.HandlerFunc
	UpdateUser() echo.HandlerFunc
	GetAllUser() echo.HandlerFunc
}

type UserUseCase interface {
	AddUser(newUser User) (User, error)
	LoginUser(userLogin User) (row int, data User, err error)
	GetProfile(id int) (User, error)
	DeleteUser(id int) (row int, err error)
	UpdateUser(id int, updateProfile User) (User, error)
	GetAllU() ([]User, error)
}

type UserData interface {
	Insert(newUser User) (User, error)
	Login(userLogin User) (row int, data User, err error)
	GetSpecific(userID int) (User, error)
	Delete(userID int) (row int, err error)
	Update(userID int, updatedData User) User
	GetAll() []User
	GetBalanceUser(userID int) (int, error)
	UpdateBalanceUser(userID, newBalance int) error
}
