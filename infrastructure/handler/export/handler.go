package export

import (
	"github.com/BrandokVargas/api-back-dportinsight/domain/export"
	"github.com/BrandokVargas/api-back-dportinsight/infrastructure/handler/response"
	"github.com/labstack/echo/v4"
)

type handler struct {
	useCase   export.UseCase
	responser response.ApiResponse
}

func newHandler(uc export.UseCase) handler {
	return handler{useCase: uc}
}

func (h handler) GetAllDataExport(c echo.Context) error {
	dataExports, err := h.useCase.GetAllDataExport()
	if err != nil {
		return h.responser.Error(c, "useCase.GetAllDataExport()", err)
	}
	return c.JSON(h.responser.OK(dataExports))
}

func (h handler) GetTopsTransportExportation(c echo.Context) error {
	tops, err := h.useCase.GetTopsTransportExportation()
	if err != nil {
		return h.responser.Error(c, "useCase.GetTopsTransportExportation()", err)
	}

	return c.JSON(h.responser.OK(tops))
}

func (h handler) GetLargestAmountExported(c echo.Context) error {
	topTenLargestAmountExported, err := h.useCase.GetLargestAmountExported()
	if err != nil {
		return h.responser.Error(c, "useCase.GetLargestAmountExported()", err)
	}
	return c.JSON(h.responser.OK(topTenLargestAmountExported))
}

func (h handler) GetAllDataLibertad(c echo.Context) error {
	getAllDataExportedLibertad, err := h.useCase.GetAllDataLibertad()
	if err != nil {
		return h.responser.Error(c, "useCase.GetAllDataLibertad()", err)
	}
	return c.JSON(h.responser.OK(getAllDataExportedLibertad))
}
