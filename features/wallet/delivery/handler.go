package delivery

import (
	"be11/apiclean/features/wallet"
	"be11/apiclean/utils/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type WalletDelivery struct {
	walletUsecase wallet.UsecaseInterface
}

func New(e *echo.Echo, usecase wallet.UsecaseInterface) {
	handler := &WalletDelivery{
		walletUsecase: usecase,
	}

	e.POST("/wallets", handler.Add)
}

func (delivery *WalletDelivery) Add(c echo.Context) error {
	var dataRequest WalletRequest
	errBind := c.Bind(&dataRequest)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("error bind data"))
	}
	row, err := delivery.walletUsecase.Add(toCore(dataRequest))
	if err != nil || row == 0 {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("error insert data"))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponseHelper("success"))
}
