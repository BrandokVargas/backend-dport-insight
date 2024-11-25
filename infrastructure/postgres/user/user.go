package user

import (
	"context"
	"database/sql"

	"github.com/BrandokVargas/api-back-dportinsight/infrastructure/postgres"
	"github.com/BrandokVargas/api-back-dportinsight/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

const table = "users"

var fields = []string{
	"id_user",
	"apellido_paterno",
	"apellido_materno",
	"email",
	"nombres",
	"dni",
	"is_admin",
	"password",
	"created_at",
	"updated_at",
}

var (
	psqlInsertUser  = postgres.BuildSQLInsert(table, fields)
	psqlGetAllUsers = postgres.BuildSQLSelect(table, fields)
)

type User struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) User {
	return User{db: db}
}

func (u User) RegisterUser(m *model.User) error {
	_, err := u.db.Exec(
		context.Background(),
		psqlInsertUser,
		m.ID,
		m.PaternalSurname,
		m.MaternalSurname,
		m.Email,
		m.Names,
		m.Dni,
		m.IsAdmin,
		m.Password,
		m.CreatedAt,
		postgres.Int64ToNull(m.UpdatedAt),
	)
	if err != nil {
		return err
	}
	return nil
}

func (u User) GetAllUsers() (model.Users, error) {
	query := psqlGetAllUsers

	rows, err := u.db.Query(
		context.Background(),
		query,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	ms := model.Users{}
	for rows.Next() {
		m, err := u.scanRow(rows, false)
		if err != nil {
			return nil, err
		}
		ms = append(ms, m)
	}

	return ms, nil
}

func (u User) scanRow(s pgx.Row, withPassword bool) (model.User, error) {
	m := model.User{}

	updatedNull := sql.NullInt64{}

	err := s.Scan(
		&m.ID,
		&m.PaternalSurname,
		&m.MaternalSurname,
		&m.Email,
		&m.Names,
		&m.Dni,
		&m.IsAdmin,
		&m.Password,
		&m.CreatedAt,
		&updatedNull,
	)
	if err != nil {
		return m, err
	}

	m.UpdatedAt = updatedNull.Int64

	if !withPassword {
		m.Password = ""
	}

	return m, nil
}
