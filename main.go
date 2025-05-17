package main

import (
	"os"

	"github.com/labstack/echo/v4"

	"fib_api/interfaces/handler"
	"fib_api/usecase"
)

func main() {
	//ポート番号を取得
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	//依存性の注入
	fibIdxUsecase := usecase.NewFibIdxUsecase()
	fibIdxHandler := handler.NewFibIdxHandler(fibIdxUsecase)

	//サーバーの初期化
	e := echo.New()
	//ルーティングの設定
	//GET /fib?n=10
	e.GET("/fib", fibIdxHandler.HandlerCalcFibNum)
	//サーバー起動
	e.Logger.Fatal(e.Start(":" + port))
}
