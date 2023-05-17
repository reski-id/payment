package factory

import (
	ud "portal/feature/user/data"
	userDelivery "portal/feature/user/delivery"
	us "portal/feature/user/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	nd "portal/feature/topup/data"
	topupDelivery "portal/feature/topup/delivery"
	nu "portal/feature/topup/usecase"

	td "portal/feature/transfer/data"
	transferDelivery "portal/feature/transfer/delivery"
	tu "portal/feature/transfer/usecase"

	pd "portal/feature/payment/data"
	paymentDelivery "portal/feature/payment/delivery"
	pu "portal/feature/payment/usecase"
)

func Initfactory(e *echo.Echo, db *gorm.DB) {
	userData := ud.New(db)
	validator := validator.New()
	useCase := us.New(userData, validator)
	userHandler := userDelivery.New(useCase)
	userDelivery.RouteUser(e, userHandler)

	topupData := nd.New(db)
	topupCase := nu.New(topupData)
	topupHandler := topupDelivery.NewTopupHandler(topupCase, userData)
	topupDelivery.RouteTopup(e, topupHandler)

	transferData := td.New(db)
	transferCase := tu.New(transferData)
	transferHandler := transferDelivery.New(transferCase)
	transferDelivery.RouteTransfer(e, transferHandler)

	paymentData := pd.New(db)
	paymentCase := pu.New(paymentData)
	paymentHandler := paymentDelivery.NewPaymentHandler(paymentCase, userData)
	paymentDelivery.RoutePayment(e, paymentHandler)

}
