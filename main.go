package main

import (
	"os"

	"github.com/labstack/echo/v4"

	"fib_api/interfaces/handler"
	"fib_api/usecase"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fibIdxUsecase := usecase.NewFibIdxUsecase()
	fibIdxHandler := handler.NewFibIdxHandler(fibIdxUsecase)

	e := echo.New()
	e.Debug = true

	e.GET("/fib", fibIdxHandler.HandlerCalcFibNum)

	e.Logger.Fatal(e.Start(":" + port))
}
