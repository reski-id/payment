package domain

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Topup struct {
	ID            int
	UserID        int
	Amount        int
	Remarks       string
	Balance       int
	BalanceBefore int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
}

type TopupHandler interface {
	InsertTopup() echo.HandlerFunc
}

type TopupUseCase interface {
	AddTopup(IDUser int, useTopup Topup) (Topup, error)
}

type TopupData interface {
	Insert(insertTopup Topup) Topup
}
