package handler

import (
	"net/http"
	"time"

	"github.com/BrandokVargas/api-back-dportinsight/infrastructure/handler/user"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, dbDport *pgxpool.Pool) {
	health(e)
	//Table user
	user.NewRouter(e, dbDport)
}

func health(e *echo.Echo) {
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(
			http.StatusOK,
			map[string]string{
				"time":         time.Now().String(),
				"message":      "Hello word",
				"service_name": "",
			},
		)
	})
}
