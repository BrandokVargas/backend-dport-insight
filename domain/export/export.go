package export

import "github.com/BrandokVargas/api-back-dportinsight/model"

type UseCase interface {
	GetAllDataExport() (model.Exports, error)
	GetTopsTransportExportation() (model.TopsTransports, error)
	GetLargestAmountExported() (model.LargestAmountExporteds, error)

	GetAllDataLibertad() (model.TotalAllDataLibertad, error)
}
type Repository interface {
	GetAllDataExport() (model.Exports, error)
	GetTopsTransportExportation() (model.TopsTransports, error)
	GetLargestAmountExported() (model.LargestAmountExporteds, error)
	GetAllDataLibertad() (model.TotalAllDataLibertad, error)
}
