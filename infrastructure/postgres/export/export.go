package export

import (
	"context"

	"github.com/BrandokVargas/api-back-dportinsight/infrastructure/postgres"
	"github.com/BrandokVargas/api-back-dportinsight/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

const table = "hecho_exportaciones"

var fields = []string{
	"exportacion_id",
	"codigo_aduana",
	"fecha_embarque",
	"codigo_agente",
	"numero_doc_exportador",
	"codigo_pais",
	"codigo_puerto",
	"codigo_cviatra",
	"codigo_transporte",
	"codigo_almacen",
	"codigo_nandina",
	"monto_exportacion",
	"peso_neto",
	"peso_bruto",
	"cantidad_exportada",
	"unidad_medidad",
	"cantidad_unidad_comercial",
	"codigo_unidad",
	"codigo_ubigeo",
}
var (
	psqlGetAllDatam                 = postgres.BuildSQLSelect(table, fields)
	psqlGetTopsTransportExportation = `SELECT vt.nombre_via AS transporte, COUNT(hc.codigo_cviatra) AS total
	FROM dim_via_transporte vt
	INNER JOIN hecho_exportaciones hc
	ON vt.codigo_via = hc.codigo_cviatra
	GROUP BY vt.nombre_via
	ORDER BY total DESC
	LIMIT 3;`

	psqlGetTopTenLargestAmountExportedToLibertad = `SELECT 
    u.provincia, ROUND(SUM(h.cantidad_exportada),0) AS cantidad_exportada, u.coordenada_a, u.coordenada_b 
		FROM hecho_exportaciones h 
		INNER JOIN dim_ubigeo u 
		ON h.codigo_ubigeo = u.codigo_ubigeo
		WHERE u.departamento = 'LA LIBERTAD'
		GROUP BY u.provincia,u.coordenada_a,u.coordenada_b
		ORDER BY cantidad_exportada DESC
		LIMIT 10;`

	psqlGetAllDataExportedLibertad = `SELECT ROUND(SUM(h.cantidad_exportada), 0) AS total_cantidad_exportada
		FROM hecho_exportaciones h INNER JOIN dim_ubigeo u ON h.codigo_ubigeo = u.codigo_ubigeo 
		WHERE u.departamento = 'LA LIBERTAD';`
)

type Export struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) Export {
	return Export{db: db}
}

func (e Export) GetAllDataExport() (model.Exports, error) {
	query := psqlGetAllDatam + " LIMIT 10"

	rows, err := e.db.Query(
		context.Background(),
		query,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	ms := model.Exports{}
	for rows.Next() {
		m, err := e.scanRow(rows)
		if err != nil {
			return nil, err
		}
		ms = append(ms, m)
	}

	return ms, nil
}
func (e Export) scanRow(s pgx.Row) (model.Export, error) {
	m := model.Export{}

	err := s.Scan(
		&m.Exportacion_id,
		&m.Codigo_aduana,
		&m.Fecha_embarque,
		&m.Codigo_agente,
		&m.Numero_doc_exportador,
		&m.Codigo_pais,
		&m.Codigo_puerto,
		&m.Codigo_cviatra,
		&m.Codigo_transporte,
		&m.Codigo_almacen,
		&m.Codigo_nandina,
		&m.Monto_exportacion,
		&m.Peso_neto,
		&m.Peso_bruto,
		&m.Cantidad_exportada,
		&m.Unidad_medidad,
		&m.Cantidad_unidad_comercial,
		&m.Codigo_unidad,
		&m.Codigo_ubigeo,
	)
	if err != nil {
		return m, err
	}
	return m, nil
}
func (e Export) GetTopsTransportExportation() (model.TopsTransports, error) {
	rows, err := e.db.Query(
		context.Background(),
		psqlGetTopsTransportExportation,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ts := model.TopsTransports{}
	for rows.Next() {
		t, err := e.scanTopsTransports(rows)
		if err != nil {
			return nil, err
		}
		ts = append(ts, t)
	}
	return ts, nil
}
func (e Export) scanTopsTransports(s pgx.Row) (model.TopTransport, error) {
	m := model.TopTransport{}
	err := s.Scan(
		&m.Transporte,
		&m.Total,
	)
	if err != nil {
		return m, err
	}
	return m, nil
}

func (e Export) GetLargestAmountExported() (model.LargestAmountExporteds, error) {
	rows, err := e.db.Query(
		context.Background(),
		psqlGetTopTenLargestAmountExportedToLibertad,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ts := model.LargestAmountExporteds{}
	for rows.Next() {
		t, err := e.scanLargestAmountExported(rows)
		if err != nil {
			return nil, err
		}
		ts = append(ts, t)
	}
	return ts, nil
}

func (e Export) scanLargestAmountExported(s pgx.Row) (model.LargestAmountExported, error) {
	m := model.LargestAmountExported{}
	err := s.Scan(
		&m.Provincia,
		&m.Cantidad_exportada,
		&m.Coordenada_a,
		&m.Coordenada_b,
	)
	if err != nil {
		return m, err
	}
	return m, nil
}

func (e Export) GetAllDataLibertad() (model.TotalAllDataLibertad, error) {

	rows, err := e.db.Query(
		context.Background(),
		psqlGetAllDataExportedLibertad,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ts := model.TotalAllDataLibertad{}
	for rows.Next() {
		t, err := e.scanAllDataLibertad(rows)
		if err != nil {
			return nil, err
		}
		ts = append(ts, t)
	}
	return ts, nil
}

func (e Export) scanAllDataLibertad(s pgx.Row) (model.TotalDataLibertad, error) {
	m := model.TotalDataLibertad{}
	err := s.Scan(
		&m.Total,
	)
	if err != nil {
		return m, err
	}
	return m, nil
}
