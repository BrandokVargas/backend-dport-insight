package export

import (
	"fmt"

	"github.com/BrandokVargas/api-back-dportinsight/model"
)

type Export struct {
	repository Repository
}

func NewExport(repository Repository) *Export {
	return &Export{repository: repository}
}

func (e Export) GetAllDataExport() (model.Exports, error) {
	dataExports, err := e.repository.GetAllDataExport()

	if err != nil {
		return nil, fmt.Errorf("%s %w ", "repository.GetAllDataExport()", err)
	}
	return dataExports, nil
}

func (e Export) GetTopsTransportExportation() (model.TopsTransports, error) {
	tops, err := e.repository.GetTopsTransportExportation()

	if err != nil {
		return nil, fmt.Errorf("%s %w", "repository.GetTopsTransportExportation()", err)
	}
	return tops, nil
}

func (e Export) GetLargestAmountExported() (model.LargestAmountExporteds, error) {
	dataLargestExportedToLibertad, err := e.repository.GetLargestAmountExported()

	if err != nil {
		return nil, fmt.Errorf("%s %w", "repository.GetLargestAmountExported()", err)
	}

	return dataLargestExportedToLibertad, nil
}

func (e Export) GetAllDataLibertad() (model.TotalAllDataLibertad, error) {
	dataAllAmountLibertad, err := e.repository.GetAllDataLibertad()
	if err != nil {
		return nil, fmt.Errorf("%s %w", "repository.GetAllDataLibertad()", err)
	}
	return dataAllAmountLibertad, nil
}
