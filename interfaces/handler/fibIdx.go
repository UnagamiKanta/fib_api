package handler

import (
	"errors"
	"fib_api/usecase"
	"net/http"
	"strconv"

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
			return c.JSON(http.StatusBadRequest, echo.Map{
				"status":  http.StatusBadRequest,
				"message": "n must be a non-negative integer",
			})
		}

		if errors.Is(err, usecase.ErrTooLargeInput) { //入力が大きすぎる場合
			return c.JSON(http.StatusBadRequest, echo.Map{
				"status":  http.StatusBadRequest,
				"message": "n is too large, please use less than " + strconv.Itoa(usecase.MAXINPUTNUM),
			})
		}
		//想定していないエラーの場合
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  http.StatusInternalServerError,
			"message": "failed to calculate Fibonacci number",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{"result": fibNum})
}
