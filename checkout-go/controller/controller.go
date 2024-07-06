package controller

import (
	"checkout/model"
	"checkout/usecase"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CheckoutController struct {
	checkoutUsecase usecase.CheckoutUsecase
}

func NewCheckoutController(usecase usecase.CheckoutUsecase) CheckoutController {
	return CheckoutController{
		checkoutUsecase: usecase,
	}
}

func (c *CheckoutController) Checkout(ctx *gin.Context) {
	var order model.Order
	err := ctx.BindJSON(&order)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	checkout, err := c.checkoutUsecase.Checkout(order)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := json.RawMessage(checkout)

	ctx.JSON(http.StatusAccepted, response)
}
