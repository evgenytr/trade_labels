package main

import (
	"github.com/labstack/echo/v4"

	"github.com/evgenytr/trade_labels.git/internal/handlers"
)

func main() {
	e := echo.New()

	e.GET("/ping", handlers.PingHandler)
	e.GET("/ping_indesign", handlers.IndesignPingHandler)
	e.POST("/orders", handlers.PostOrderHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
