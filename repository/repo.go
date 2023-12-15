package repository

import (
	"context"
	"database/sql"
	"ginlearning/model/domain"
)

type Repository interface {
	Login(ctx context.Context, tx *sql.Tx, LoginCategory domain.LoginData) (domain.LoginData, error)
	Register(ctx context.Context, tx *sql.Tx, RegisterCategory domain.RegisterData)
}
