package repository

import (
	"context"
	"database/sql"
	"task2/model/domain"
)

type UserRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) []domain.User
	FindById(ctx context.Context, tx *sql.Tx, id int) domain.User
	FindByUsername(ctx context.Context, tx *sql.Tx, username string) domain.User
	Login(ctx context.Context, tx *sql.Tx, user domain.User) error
	Register(ctx context.Context, tx *sql.Tx, user domain.User)
	UpdatePoint(ctx context.Context, tx *sql.Tx, user domain.User)
}
