package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"ginlearning/model/domain"
	"ginlearning/security"
)

type RepoImpl struct {
}

func NewRepository() Repository {
	return &RepoImpl{}
}

func (repo *RepoImpl) Login(ctx context.Context, tx *sql.Tx, LoginCategory domain.LoginData) (domain.LoginData, error) {
	sql := "SELECT gmail, passwords FROM users WHERE gmail = ?"
	rows, err := tx.QueryContext(ctx, sql, LoginCategory.Gmail)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	loglist := domain.LoginData{}
	if rows.Next() {
		err := rows.Scan(&loglist.Gmail, &loglist.Passwords)
		if err != nil {
			panic(err)
		}
		token, eror := security.ClaimsJwt(loglist, LoginCategory.Passwords)
		if eror != nil {
			panic("error jwt")
		}
		fmt.Println(token)
		loglist.Token = token
		return loglist, nil
	} else {
		return loglist, errors.New("data not found")
	}
}

func (repo *RepoImpl) Register(ctx context.Context, tx *sql.Tx, RegisterCategory domain.RegisterData) {
	sql := "insert into users(gmail,names,passwords) values (?,?,?)"
	encrypt, err := security.PassEncrypt(RegisterCategory)
	if err != nil {
		panic("error from hashing")
	}
	result, err := tx.ExecContext(ctx, sql, RegisterCategory.Gmail, RegisterCategory.Names, encrypt)
	if err != nil {
		panic(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	RegisterCategory.Id = int(id)
}
