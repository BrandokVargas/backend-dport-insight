package export

import (
	"github.com/BrandokVargas/api-back-dportinsight/domain/export"
	repostitoryExport "github.com/BrandokVargas/api-back-dportinsight/infrastructure/postgres/export"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

func NewRouter(e *echo.Echo, db *pgxpool.Pool) {
	h := buildHandler(db)
	publicRoutes(e, h)
}

func buildHandler(dbPool *pgxpool.Pool) handler {

	repository := repostitoryExport.New(dbPool)
	useCase := export.NewExport(repository)

	return newHandler(useCase)
}

// ROUTES PUBLIC
func publicRoutes(e *echo.Echo, h handler) {
	g := e.Group("api/v1/dataexports")
	g.GET("", h.GetAllDataExport)

	getTops := e.Group("api/v1/topsTransports")
	getTops.GET("", h.GetTopsTransportExportation)

	getTopTenLargestExported := e.Group("api/v1/topTenLargestAmountExported")
	getTopTenLargestExported.GET("", h.GetLargestAmountExported)

	getAmountDataExporteLibertad := e.Group("api/v1/getAmountAllDataLibertad")
	getAmountDataExporteLibertad.GET("", h.GetAllDataLibertad)

}
