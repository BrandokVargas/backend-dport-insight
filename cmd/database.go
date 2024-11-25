package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	// Tiempo máximo de espera para las conexiones
	dbTimeout    = 5 * time.Second
	AppName      = "DportInsight"
	AppNameDport = "DportInsightDb"
)

// Crear conexión a la primera base de datos
func NewDBConnection() (*pgxpool.Pool, error) {
	// Definir número de conexiones mínimas y máximas
	min := 3
	max := 100

	// Obtener variables de entorno
	minConnection := os.Getenv("DB_MIN_CONN")
	maxConnection := os.Getenv("DB_MAX_CONN")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	sslMode := os.Getenv("DB_SSL_MODE")

	// Validar y ajustar las conexiones
	if minConnection != "" {
		v, err := strconv.Atoi(minConnection)
		if err != nil {
			log.Println("Warning: DB_MIN_CONN has not a valid value, defaulting to", min)
		} else if v >= min && v <= max {
			min = v
		}
	}

	if maxConnection != "" {
		v, err := strconv.Atoi(maxConnection)
		if err != nil {
			log.Println("Warning: DB_MAX_CONN has not a valid value, defaulting to", max)
		} else if v >= min && v <= max {
			max = v
		}
	}

	// Crear cadena de conexión
	connString := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s pool_min_conns=%d pool_max_conns=%d",
		user, pass, host, port, dbName, sslMode, min, max,
	)

	// Crear el config para la base de datos
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("pgxpool.ParseConfig() %w", err)
	}

	config.ConnConfig.RuntimeParams["application_name"] = AppName

	// Crear contexto con timeout
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// Crear pool de conexiones
	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("pgxpool.NewWithConfig() %w", err)
	}

	// Verificar la conexión con ping
	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("unable to ping database: %w", err)
	}

	return pool, nil
}

// Crear conexión para la segunda base de datos
func NewDBConnectionDport() (*pgxpool.Pool, error) {
	min := 3
	max := 100

	// Obtener variables de entorno
	minConnection := os.Getenv("DB_MIN_CONN")
	maxConnection := os.Getenv("DB_MAX_CONN")
	user := os.Getenv("DB_USER_DB_INTEGRATED")
	pass := os.Getenv("DB_PASSWORD_DB_INTEGRATED")
	host := os.Getenv("DB_HOST_DB_INTEGRATED")
	port := os.Getenv("DB_PORT_DB_INTEGRATED")
	dbName := os.Getenv("DB_NAME_DB_INTEGRATED")
	sslMode := os.Getenv("DB_SSL_MODE_DB_INTEGRATED")

	// Validar y ajustar las conexiones
	if minConnection != "" {
		v, err := strconv.Atoi(minConnection)
		if err != nil {
			log.Println("Warning: DB_MIN_CONN has not a valid value, defaulting to", min)
		} else if v >= min && v <= max {
			min = v
		}
	}

	if maxConnection != "" {
		v, err := strconv.Atoi(maxConnection)
		if err != nil {
			log.Println("Warning: DB_MAX_CONN has not a valid value, defaulting to", max)
		} else if v >= min && v <= max {
			max = v
		}
	}

	// Crear cadena de conexión
	connString := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s pool_min_conns=%d pool_max_conns=%d",
		user, pass, host, port, dbName, sslMode, min, max,
	)

	// Crear el config para la base de datos
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("pgxpool.ParseConfig() %w", err)
	}

	config.ConnConfig.RuntimeParams["application_name"] = AppNameDport

	// Crear contexto con timeout
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// Crear pool de conexiones
	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("pgxpool.NewWithConfig() %w", err)
	}

	// Verificar la conexión con ping
	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("unable to ping database: %w", err)
	}

	return pool, nil
}
