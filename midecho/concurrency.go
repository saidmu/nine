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
		defer func() { <-c.Queue }()
		return func(c echo.Context) error {
			return next(c)
		}
	default:
		return func(c echo.Context) error {
			c.JSON(http.StatusTooManyRequests,echo.Map{"Message": "Too many requests"})
			return nil
		}
	}
}