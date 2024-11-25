package model

type LargestAmountExported struct {
	Provincia          string `json:"provincia"`
	Cantidad_exportada string `json:"cantidad_exportada"`
	Coordenada_a       string `json:"coordenada_a"`
	Coordenada_b       string `json:"coordenada_b"`
}

type LargestAmountExporteds []LargestAmountExported
