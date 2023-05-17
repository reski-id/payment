package delivery

import (
	"fmt"
	"log"
	"net/http"
	"portal/domain"
	"portal/feature/common"

	"github.com/labstack/echo/v4"
)

type transferHandler struct {
	transferUseCase domain.TransferUseCase
}

func New(nu domain.TransferUseCase) domain.TransferHandler {
	return &transferHandler{
		transferUseCase: nu,
	}
}

func (nh *transferHandler) InsertTransfer() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmp InsertRequest
		err := c.Bind(&tmp)

		if err != nil {
			log.Println("Cannot parse data", err)
			c.JSON(http.StatusBadRequest, "error read input")
		}

		fmt.Println(tmp)

		var userid, _ = common.ExtractData(c)
		data, err := nh.transferUseCase.AddTransfer(userid, tmp.ToDomain())

		if err != nil {
			log.Println("Cannot proces data", err)
			c.JSON(http.StatusInternalServerError, err)
		}

		fmt.Println(userid)

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "success create data",
			"data":    FromDomain(data),
		})

	}
}
