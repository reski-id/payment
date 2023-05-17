package delivery

import (
	"fmt"
	"log"
	"net/http"
	"portal/domain"
	"portal/feature/common"

	"github.com/labstack/echo/v4"
)

type topupHandler struct {
	topupUseCase domain.TopupUseCase
	userData     domain.UserData
}

func NewTopupHandler(useCase domain.TopupUseCase, userData domain.UserData) domain.TopupHandler {
	return &topupHandler{
		topupUseCase: useCase,
		userData:     userData,
	}
}

func (nh *topupHandler) InsertTopup() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmp InsertRequest
		err := c.Bind(&tmp)
		if err != nil {
			log.Println("Cannot parse data", err)
			return c.JSON(http.StatusBadRequest, "error read input")
		}

		userID, _ := common.ExtractData(c)

		// Step 1: Check user balance
		userBalance, err := nh.userData.GetBalanceUser(userID)
		if err != nil {
			log.Println("Error retrieving user balance:", err)
			return c.JSON(http.StatusInternalServerError, "error retrieving user balance")
		}

		newTopup := tmp.ToDomain()
		newTopup.UserID = userID

		// Step 2: Update user balance with new top-up amount
		newBalance := userBalance + newTopup.Amount
		err = nh.userData.UpdateBalanceUser(userID, newBalance)
		if err != nil {
			log.Println("Error updating user balance:", err)
			return c.JSON(http.StatusInternalServerError, "error updating user balance")
		}

		// Step 3: Save top-up data in the topup table
		data, err := nh.topupUseCase.AddTopup(userID, newTopup)
		if err != nil {
			log.Println("Cannot process data", err)
			return c.JSON(http.StatusInternalServerError, "error processing top-up data")
		}

		// Step 4: Save log in the userlog table
		logData := domain.UserLog{
			Balance:         newBalance,
			BalanceBefore:   userBalance,
			TransactionType: 1, // Assuming 1 represents a top-up transaction
			TrannsctionLog:  "Top-up transaction",
		}
		// _, err = nh.userLogData.Insert(logData)
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
