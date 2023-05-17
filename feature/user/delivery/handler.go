package delivery

import (
	"log"
	"net/http"
	"portal/domain"
	"portal/feature/common"
	"strings"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userUsecase domain.UserUseCase
}

func New(nu domain.UserUseCase) domain.UserHandler {
	return &userHandler{
		userUsecase: nu,
	}
}

func (uh *userHandler) InsertUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmp InsertFormat
		err := c.Bind(&tmp)

		if err != nil {
			log.Println("Cannot parse data", err)
			c.JSON(http.StatusBadRequest, "error read input")
		}

		data, err := uh.userUsecase.AddUser(tmp.ToModel())

		if err != nil {
			log.Println("Cannot proces data", err)
			c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "success create data",
			"data":    data,
			"token":   common.GenerateToken(data.ID, data.PhoneNumber),
		})
	}
}

func (uh *userHandler) LoginUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var userLogin LoginFormat
		err := c.Bind(&userLogin)
		if err != nil {
			log.Println("Cannot parse data", err)
			return c.JSON(http.StatusBadRequest, "cannot read input")
		}
		row, data, e := uh.userUsecase.LoginUser(userLogin.LoginToModel())
		if e != nil {
			log.Println("Cannot proces data", err)
			return c.JSON(http.StatusBadRequest, "Phone  Number or PIN incorrect")
		}
		if row == -1 {
			return c.JSON(http.StatusBadRequest, "Phone Number or PIN incorrect")
		}

		token := common.GenerateToken(int(data.ID), data.PhoneNumber)

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success login",
			"token":   token,
			"data":    data,
		})
	}
}

func (uh *userHandler) GetProfile() echo.HandlerFunc {
	return func(c echo.Context) error {

		id, _ := common.ExtractData(c)

		data, err := uh.userUsecase.GetProfile(id)

		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusNotFound, err.Error())
			} else {
				return c.JSON(http.StatusInternalServerError, err.Error())
			}
		}
		return c.JSON(http.StatusFound, map[string]interface{}{
			"message": "data found",
			"data":    data,
		})
	}
}

func (uh *userHandler) DeleteUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := common.ExtractData(c)
		if id == 0 {
			return c.JSON(http.StatusUnauthorized, "Unauthorized")
		}
		_, errDel := uh.userUsecase.DeleteUser(id)
		if errDel != nil {
			return c.JSON(http.StatusInternalServerError, "cannot delete user")
		}
		return c.JSON(http.StatusOK, "success to delete user")
	}
}

func (uh *userHandler) UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {

		var tmp InsertFormat
		result := c.Bind(&tmp)

		idUpdate, _ := common.ExtractData(c)

		if result != nil {
			log.Println(result, "Cannot parse input to object")
			return c.JSON(http.StatusInternalServerError, "Error read update")
		}

		_, err := uh.userUsecase.UpdateUser(idUpdate, tmp.ToModel())

		if err != nil {
			return c.JSON(http.StatusInternalServerError, "cannot update")
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "Success update data",
		})
	}
}

func (uh *userHandler) GetAllUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, role := common.ExtractData(c)

		if role != "admin" {
			return c.JSON(http.StatusCreated, map[string]interface{}{
				"message": "Only Admin can access",
			})
		}
		data, err := uh.userUsecase.GetAllU()
		if err != nil {
			log.Println("Cannot get data", err)
			return c.JSON(http.StatusBadRequest, "error read input")

		}

		if data == nil {
			log.Println("Terdapat error saat mengambil data")
			return c.JSON(http.StatusInternalServerError, "Problem from database")
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get all data",
			"data":    data,
		})
	}
}
