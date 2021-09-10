package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Stats(c echo.Context) error {
	log.Println("loading stats...")

	s, err := h.repository.GetStats()
	if err != nil {
		log.Println("error while loading stats,", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return h.SendResult(c, s)
}
