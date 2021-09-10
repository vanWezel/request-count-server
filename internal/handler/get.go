package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Get(c echo.Context) error {
	log.Println("adding to stats...")

	s, err := h.repository.Increment()
	if err != nil {
		log.Println("error while adding to stats,", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return h.SendResult(c, s)
}
