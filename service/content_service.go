package service

import (
	"context"
	"task2/model/request"
	"task2/model/response"
)

type ContentService interface {
	FindAll(ctx context.Context) ([]response.ResponseUserContentAll, error)
	FindById(ctx context.Context, id int) (response.ResponseUserContentByContentId, error)
	FindByStatusPublish(ctx context.Context) ([]response.ResponseUserContentByContentId, error)
	FindByStatusDraft(ctx context.Context) ([]response.ResponseUserContentByContentId, error)
	Create(ctx context.Context, request request.RequestCreateContent, user_id int) (response.ResponseUserContentCreated, error)
	Update(ctx context.Context, request request.RequestUpdateContent, user_id, content_id int) (response.ResponseUserContentUpdateByContentId, error)
	Delete(ctx context.Context, user_id, content_id int) error
	FindContentByUserId(ctx context.Context, user_id, content_id int) (response.ResponseUserContentByContentId, error)
}
