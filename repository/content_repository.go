package repository

import (
	"context"
	"database/sql"
	"task2/model/domain"
)

type ContentRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Content
	FindById(ctx context.Context, tx *sql.Tx, id int) domain.Content
	FindByStatusPublish(ctx context.Context, tx *sql.Tx) []domain.Content
	FindByStatusDraft(ctx context.Context, tx *sql.Tx) []domain.Content
	Create(ctx context.Context, tx *sql.Tx, content domain.Content) *int
	Update(ctx context.Context, tx *sql.Tx, content domain.Content)
	Delete(ctx context.Context, tx *sql.Tx, id int)
	FindContentByContentIdAndUserId(ctx context.Context, tx *sql.Tx, user_id, content_id int) domain.Content
}
