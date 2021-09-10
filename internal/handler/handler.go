package handler

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"

	"github.com/vanWezel/request-count-server/internal/model"
	"github.com/vanWezel/request-count-server/internal/repository"
	"github.com/vanWezel/request-count-server/pkg/helper"
)

type Handler struct {
	repository repository.Interface
	instance   string
	Port       int
}

func New() *Handler {
	host := helper.Getenv("REDIS_HOST", "redis:6379")
	db := helper.GetenvToi("REDIS_DB", "0")
	instance := helper.Getenv("HOSTNAME", "localhost")
	port := helper.GetenvToi("PORT", "80")

	return &Handler{
		repository: repository.NewRedis(host, db, instance),
		instance:   instance,
		Port:       port,
	}
}

func (h *Handler) SendResult(c echo.Context, s *model.Stats) error {
	log.Println("sending result...")

	var r []string
	r = append(r, fmt.Sprintf("You are talking to instance %v:%v.", h.instance, h.Port))
	r = append(r, fmt.Sprintf("This is request %v to this instance and request %v to the cluster.", s.Instance, s.Total))

	return c.String(http.StatusOK, strings.Join(r, "\n"))
}
