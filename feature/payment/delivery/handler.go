package delivery

import (
	"fmt"
	"log"
	"net/http"
	"portal/domain"
	"portal/feature/common"

	"github.com/labstack/echo/v4"
)

type paymentHandler struct {
	paymentUseCase domain.PaymentUseCase
	userData       domain.UserData
}

func NewPaymentHandler(paymentUseCase domain.PaymentUseCase, userData domain.UserData) domain.PaymentHandler {
	return &paymentHandler{
		paymentUseCase: paymentUseCase,
		userData:       userData,
	}
}

func (ph *paymentHandler) InsertPayment() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmp InsertRequest
		err := c.Bind(&tmp)
		if err != nil {
			log.Println("Cannot parse data", err)
			return c.JSON(http.StatusBadRequest, "error reading input")
		}

		userID, _ := common.ExtractData(c)

		// Step 1: Check user balance
		userBalance, err := ph.userData.GetBalanceUser(userID)
		if err != nil {
			log.Println("Error retrieving user balance:", err)
			return c.JSON(http.StatusInternalServerError, "error retrieving user balance")
		}

		newPayment := tmp.ToDomain()
		newPayment.UserID = userID

		// Step 2: Update user balance with the payment amount
		newBalance := userBalance - newPayment.Amount
		err = ph.userData.UpdateBalanceUser(userID, newBalance)
		if err != nil {
			log.Println("Error updating user balance:", err)
			return c.JSON(http.StatusInternalServerError, "error updating user balance")
		}

		// Step 3: Save payment data in the payment table
		newPayment.BalanceBefore = userBalance
		newPayment.Balance = newBalance
		data, err := ph.paymentUseCase.AddPayment(userID, newPayment)
		if err != nil {
			log.Println("Cannot process data", err)
			return c.JSON(http.StatusInternalServerError, "error processing payment data")
		}

		// Step 4: Save log in the userlog table
		logData := domain.UserLog{
			Balance:         newBalance,
			BalanceBefore:   userBalance,
			TransactionType: 2, // Assuming 2 represents a payment transaction
			TrannsctionLog:  "Payment transaction",
		}
		// _, err = ph.userLogData.Insert(logData)
		// if err != nil {
		// 	log.Println("Error saving user log:", err)
		// 	return c.JSON(http.StatusInternalServerError, "error saving user log")
		// }

		fmt.Println(logData)
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "success create data",
			"data":    FromDomain(data),
		})
	}
}
