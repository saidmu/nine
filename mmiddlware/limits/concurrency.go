package limits

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type ConcurConfig struct {
	Queue chan struct{}
}

func (c ConcurConfig) Limit(next echo.HandlerFunc) echo.HandlerFunc  {
	select {
	case c.Queue <- struct{}{}:
		return func(e echo.Context) error {
			return next(e)
		}
	default:
		return func(e echo.Context) error {
			return e.JSON(http.StatusTooManyRequests,echo.Map{"Message": "Too many requests"})
		}
	}
}