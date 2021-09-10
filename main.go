package main

import (
	"fmt"

	"github.com/labstack/echo/v4"

	"github.com/vanWezel/request-count-server/internal/handler"
)

func main() {
	e := echo.New()

	h := handler.New()
	e.GET("/health", h.Health)
	e.GET("/", h.Get)
	e.GET("/stats", h.Stats)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", h.Port)))
}
