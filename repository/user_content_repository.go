package repository

import (
	"context"
	"database/sql"
	"task2/model/domain"
)

type UserContentRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) []domain.UserContent
	FindAllContentByUserId(ctx context.Context, tx *sql.Tx) []domain.UserContent
	FindById(ctx context.Context, tx *sql.Tx, content_id int) domain.UserContent
	FindByContentIdAndUserId(ctx context.Context, tx *sql.Tx, user_id, content_id int) domain.UserContent
	Create(ctx context.Context, tx *sql.Tx, userContent domain.UserContent)
	UpdatePoint(ctx context.Context, tx *sql.Tx, userContent domain.UserContent)
	DeleteByUserId(ctx context.Context, tx *sql.Tx, user_id int)
	DeleteByContentId(ctx context.Context, tx *sql.Tx, user_id, content_id int)
	Delete(ctx context.Context, tx *sql.Tx, content_id int)
}
