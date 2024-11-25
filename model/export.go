package model

import "time"

type Export struct {
	Exportacion_id            string    `json:"exportacion_id"`
	Codigo_aduana             string    `json:"codigo_aduana"`
	Fecha_embarque            time.Time `json:"fecha_embarque"`
	Codigo_agente             string    `json:"codigo_agente"`
	Numero_doc_exportador     string    `json:"numero_doc_exportador"`
	Codigo_pais               string    `json:"codigo_pais"`
	Codigo_puerto             string    `json:"codigo_puerto"`
	Codigo_cviatra            string    `json:"codigo_cviatra"`
	Codigo_transporte         string    `json:"codigo_transporte"`
	Codigo_almacen            string    `json:"codigo_almacen"`
	Codigo_nandina            string    `json:"codigo_nandina"`
	Monto_exportacion         float64   `json:"monto_exportacion"`
	Peso_neto                 float64   `json:"peso_neto"`
	Peso_bruto                float64   `json:"peso_bruto"`
	Cantidad_exportada        float64   `json:"cantidad_exportada"`
	Unidad_medidad            string    `json:"unidad_medidad"`
	Cantidad_unidad_comercial string    `json:"cantidad_unidad_comercial"`
	Codigo_unidad             string    `json:"codigo_unidad"`
	Codigo_ubigeo             string    `json:"codigo_ubigeo"`
}

type Exports []Export
