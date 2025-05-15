package handler

import (
	"net/http"

	"fib_api/usecase"

	"github.com/labstack/echo/v4"
)

type FibIdxHandler interface {
	HandlerCalcFibNum(c echo.Context) error
}

type fibIdxHandler struct {
	fibIdxUsecase usecase.FibIdxUsecase
}

func NewFibIdxHandler(fu usecase.FibIdxUsecase) FibIdxHandler {
	return &fibIdxHandler{
		fibIdxUsecase: fu,
	}
}

func (fh *fibIdxHandler) HandlerCalcFibNum(c echo.Context) error {
	fibIdxStr := c.QueryParam("n")

	fibNum, err := fh.fibIdxUsecase.CalcFibNum(fibIdxStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "n must be less than or equal to 93"})
	}

	return c.JSON(http.StatusOK, echo.Map{"result": fibNum})
}
