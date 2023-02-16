package service

import (
	"context"
	"task2/model/request"
	"task2/model/response"
)

type UserContentService interface {
	PostReward(ctx context.Context, request request.RequestCreateUserContent) (response.ResponseUserContentPoint, error)
}
