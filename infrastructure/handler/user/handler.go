package user

import (
	"github.com/BrandokVargas/api-back-dportinsight/domain/user"
	"github.com/BrandokVargas/api-back-dportinsight/infrastructure/handler/response"
	"github.com/BrandokVargas/api-back-dportinsight/model"
	"github.com/labstack/echo/v4"
)

type handler struct {
	useCase   user.UseCase
	responser response.ApiResponse
}

func newHandler(uc user.UseCase) handler {
	return handler{useCase: uc}
}

func (h handler) RegisterUser(c echo.Context) error {
	m := model.User{}

	if err := c.Bind(&m); err != nil {
		// return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return h.responser.BindFailed(err)
	}

	if err := h.useCase.RegisterUser(&m); err != nil {
		// return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return h.responser.Error(c, "useCase.RegisterUser()", err)
	}

	return c.JSON(h.responser.Created(m))
}

func (h handler) GetAllUsers(c echo.Context) error {
	users, err := h.useCase.GetAllUsers()
	if err != nil {
		return h.responser.Error(c, "useCase.GetAllUsers()", err)
	}
	return c.JSON(h.responser.OK(users))
}
