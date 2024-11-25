package response

import (
	"net/http"

	"github.com/BrandokVargas/api-back-dportinsight/model"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

const (
	BindFailed      = "bind_failed"
	Ok              = "ok"
	Created         = "created"
	Updated         = "updated"
	Delete          = "delete"
	UnexpectedError = "unexpected_error"
	AuthError       = "authorization_error"
)

type ApiResponse struct{}

func New() ApiResponse {
	return ApiResponse{}
}

func (a ApiResponse) OK(data any) (int, model.MessageResponse) {
	return http.StatusOK, model.MessageResponse{
		Data:     data,
		Messages: model.Responses{{Code: Ok, Message: "¡ready!"}},
	}
}

func (a ApiResponse) Created(data any) (int, model.MessageResponse) {
	return http.StatusOK, model.MessageResponse{
		Data:     data,
		Messages: model.Responses{{Code: Created, Message: "¡ready!"}},
	}
}

func (a ApiResponse) Updated(data any) (int, model.MessageResponse) {
	return http.StatusOK, model.MessageResponse{
		Data:     data,
		Messages: model.Responses{{Code: Updated, Message: "¡ready!"}},
	}
}

func (a ApiResponse) Delete(data any) (int, model.MessageResponse) {
	return http.StatusOK, model.MessageResponse{
		Data:     data,
		Messages: model.Responses{{Code: Delete, Message: "¡ready!"}},
	}
}

func (a ApiResponse) BindFailed(err error) error {

	customError := model.NewError()
	customError.Err = err
	customError.Code = BindFailed
	customError.Status = http.StatusBadRequest
	customError.Who = "c.Bind()"

	log.Warnf("%s", customError.Error())

	return &customError
}

func (a ApiResponse) Error(c echo.Context, who string, err error) *model.Error {
	customError := model.NewError()
	customError.Err = err
	customError.Message = "¡Ups! ocurrió un problema, disculpanos lo solucionaremos"
	customError.Code = UnexpectedError
	customError.Status = http.StatusInternalServerError
	customError.Who = who

	userID, ok := c.Get("userID").(uuid.UUID)

	if !ok {
		log.Errorf("cannot get/parse uuidfrom userID")
	}
	customError.UserID = userID.String()

	log.Errorf("%s", customError.Error())
	return &customError
}
