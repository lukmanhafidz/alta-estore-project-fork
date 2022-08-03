package delivery

import (
	"altaproject2/domain"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type cartHandler struct {
	cartUseCase domain.CartUseCase
}

func New(ps domain.CartUseCase) domain.CartHandler {
	return &cartHandler{
		cartUseCase: ps,
	}
}

// Delete implements domain.CartHandler
func (*cartHandler) Delete() echo.HandlerFunc {
	panic("unimplemented")
}

// Get implements domain.CartHandler
func (*cartHandler) Get() echo.HandlerFunc {
	panic("unimplemented")
}

// Post implements domain.CartHandler
func (*cartHandler) Post() echo.HandlerFunc {
	panic("unimplemented")
}

// Update implements domain.CartHandler
func (ch *cartHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var updatecart CartFormat
		id := c.Param("productId")
		cnv, _ := strconv.Atoi(id)
		bind := c.Bind(&updatecart)

		if bind != nil {
			log.Println("cant bind")
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": "There is an error in internal server",
			})
		}

		status := ch.cartUseCase.UpdateData(updatecart.ToModel(), cnv)

		if status == 400 {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    status,
				"message": "Wrong input",
			})
		}

		if status == 500 {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    status,
				"message": "There is an error in internal server",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    status,
			"message": "Update success",
		})
	}
}
