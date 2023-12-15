package service

import (
	"context"
	"database/sql"
	"errors"
	"ginlearning/helper"
	"ginlearning/model/domain"
	"ginlearning/model/web"
	"ginlearning/repository"

	"github.com/go-playground/validator/v10"
)

type ServicesImpl struct {
	Repository repository.Repository
	Db         *sql.DB
	Validates  *validator.Validate
}

func NewServices(repo repository.Repository, Db *sql.DB, validate *validator.Validate) Service {
	return &ServicesImpl{Repository: repo, Db: Db, Validates: validate}
}

func (services *ServicesImpl) Login(ctx context.Context, request web.LoginRequest) (domain.LoginData, error) {
	err := services.Validates.Struct(request)
	if err != nil {
		panic(err)
	}
	tx, err := services.Db.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.TxErrHandle(tx)
	category := domain.LoginData{
		Gmail:     request.Gmail,
		Passwords: request.Passwords,
	}
	token, errs := services.Repository.Login(ctx, tx, category)
	if errs != nil {
		return token, errors.New("error at service pattern")
	}
	return token, nil
}

func (services *ServicesImpl) Register(ctx context.Context, request web.RegisterRequest) {
	err := services.Validates.Struct(request)
	helper.PanicErrorIf(err)
	tx, err := services.Db.Begin() //open database
	if err != nil {
		panic(err)
	}
	defer helper.TxErrHandle(tx)
	category := domain.RegisterData{
		Gmail:     request.Gmail,
		Names:     request.Names,
		Passwords: request.Passwords,
	}

	services.Repository.Register(ctx, tx, category)
}
