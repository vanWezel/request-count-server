package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Health(c echo.Context) error {
	if err := h.repository.Ping(); err != nil {
		log.Println("error while executing ping,", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, "Yup, still working!")
}
