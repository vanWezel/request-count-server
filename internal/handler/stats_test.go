
package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"github.com/vanWezel/request-count-server/internal/repository"
)

func TestStatsSuccess(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/stats", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/stats")

	replica := "test"
	port := 80
	h := &Handler{repository.NewMock(), replica, port}

	var r []string
	r = append(r, fmt.Sprintf("You are talking to instance %v:%v.", replica, port))
	r = append(r, fmt.Sprintf("This is request %v to this instance and request %v to the cluster.", 1, 1))
	expectedBody := strings.Join(r, "\n")

	if assert.NoError(t, h.Get(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expectedBody, rec.Body.String())
	}
}
