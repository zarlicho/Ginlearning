package service

import (
	"context"
	"ginlearning/model/domain"
	"ginlearning/model/web"
)

type Service interface {
	Login(ctx context.Context, requests web.LoginRequest) (domain.LoginData, error)
	Register(ctx context.Context, requests web.RegisterRequest)
}
