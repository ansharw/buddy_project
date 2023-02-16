package service

import (
	"context"
	"task2/model/request"
	"task2/model/response"
)

type UserService interface {
	FindAllRegisteredUser(ctx context.Context) ([]response.ResponseRegisteredUser, error)
	FindAllUserPoint(ctx context.Context) ([]response.ResponseRewardPoint, error)
	FindPointByUserId(ctx context.Context, user_id int) (response.ResponseUserPoint, error)
	Login(ctx context.Context, request request.RequestLoginUser) error
	Register(ctx context.Context, request request.RequestRegisterUser) error
}
