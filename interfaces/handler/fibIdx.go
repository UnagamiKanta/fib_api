package handler

import (
	"errors"
	"fib_api/usecase"
	"net/http"

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
		if errors.Is(err, usecase.ErrInvalidInput) { //不正な入力の場合
			return c.JSON(http.StatusBadRequest, echo.Map{"message": "n must be a non-negative integer"})
		}
		//想定していないエラーの場合
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "failed to calculate Fibonacci number"})
	}

	return c.JSON(http.StatusOK, echo.Map{"result": fibNum})
}
