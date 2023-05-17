package domain

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Payment struct {
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

type PaymentHandler interface {
	InsertPayment() echo.HandlerFunc
}

type PaymentUseCase interface {
	AddPayment(IDUser int, usePayment Payment) (Payment, error)
}

type PaymentData interface {
	Insert(insertPayment Payment) Payment
}
