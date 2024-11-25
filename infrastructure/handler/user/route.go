package user

import (
	"github.com/BrandokVargas/api-back-dportinsight/domain/user"
	repositoryUser "github.com/BrandokVargas/api-back-dportinsight/infrastructure/postgres/user"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

func NewRouter(e *echo.Echo, dbPool *pgxpool.Pool) {
	h := buildHandler(dbPool)
	publicRoutes(e, h)
}

func buildHandler(dbPool *pgxpool.Pool) handler {
	repository := repositoryUser.New(dbPool)
	useCase := user.NewUser(repository)

	return newHandler(useCase)
}

func publicRoutes(e *echo.Echo, h handler) {

	//Router register User
	group := e.Group("/api/v1/public/users")
	group.POST("", h.RegisterUser)
}
