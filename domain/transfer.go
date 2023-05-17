package domain

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Transfer struct {
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

type TransferHandler interface {
	InsertTransfer() echo.HandlerFunc
}

type TransferUseCase interface {
	AddTransfer(IDUser int, useTransfer Transfer) (Transfer, error)
}

type TransferData interface {
	Insert(insertTransfer Transfer) Transfer
}
