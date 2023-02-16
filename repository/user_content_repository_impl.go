package repository

import (
	"context"
	"database/sql"
	"task2/model/domain"
)

type userContentRepository struct {
}

func NewUserContentRepository() *userContentRepository {
	return &userContentRepository{}
}

func (repo *userContentRepository) FindAll(ctx context.Context, tx *sql.Tx) []domain.UserContent {
	query := "SELECT user_id, content_id, point FROM user_content ORDER BY content_id DESC"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	var userContents []domain.UserContent
	for rows.Next() {
		var userContent domain.UserContent
		rows.Scan(userContent.GetUserId(), userContent.GetContentId(), userContent.GetPoint())
		userContents = append(userContents, userContent)
	}
	return userContents
}

func (repo *userContentRepository) FindAllContentByUserId(ctx context.Context, tx *sql.Tx) []domain.UserContent {
	query := "SELECT user_id, content_id FROM user_content WHERE user_id = ? ORDER BY content_id DESC"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	var userContents []domain.UserContent
	for rows.Next() {
		var userContent domain.UserContent
		rows.Scan(userContent.GetUserId(), userContent.GetContentId())
		userContents = append(userContents, userContent)
	}
	return userContents
}

func (repo *userContentRepository) FindById(ctx context.Context, tx *sql.Tx, content_id int) domain.UserContent {
	query := "SELECT user_id, content_id FROM user_content WHERE content_id = ?"
	userContent := domain.UserContent{}
	row := tx.QueryRowContext(ctx, query, content_id)
	row.Scan(userContent.GetContentId())
	return userContent
}

func (repo *userContentRepository) FindByContentIdAndUserId(ctx context.Context, tx *sql.Tx, user_id, content_id int) domain.UserContent {
	query := "SELECT user_id, content_id FROM user_content WHERE user_id = ? AND content_id = ? ORDER BY content_id DESC"
	userContent := domain.UserContent{}
	row := tx.QueryRowContext(ctx, query, user_id, content_id)
	row.Scan(userContent.GetUserId(), userContent.GetContentId())
	return userContent
}

func (repo *userContentRepository) Create(ctx context.Context, tx *sql.Tx, userContent domain.UserContent) {
	query := "INSERT INTO user_content(user_id, content_id) VALUES(?, ?)"
	_, err := tx.ExecContext(ctx, query, userContent.GetUserId(), userContent.GetContentId())
	if err != nil {
		panic(err)
	}
}

func (repo *userContentRepository) UpdatePoint(ctx context.Context, tx *sql.Tx, userContent domain.UserContent) {
	query := "UPDATE user_content SET point = ? WHERE user_id = ? AND content_id = ?"
	_, err := tx.ExecContext(ctx, query, userContent.GetPoint(), userContent.GetUserId(), userContent.GetContentId())
	if err != nil {
		panic(err)
	}
}

func (repo *userContentRepository) DeleteByUserId(ctx context.Context, tx *sql.Tx, user_id int) {
	query := "DELETE FROM user_content WHERE user_id = ?"
	_, err := tx.ExecContext(ctx, query, user_id)
	if err != nil {
		panic(err)
	}
}

func (repo *userContentRepository) DeleteByContentId(ctx context.Context, tx *sql.Tx, user_id, content_id int) {
	query := "DELETE FROM user_content WHERE user_id = ? AND content_id = ?"
	_, err := tx.ExecContext(ctx, query, user_id, content_id)
	if err != nil {
		panic(err)
	}
}

func (repo *userContentRepository) Delete(ctx context.Context, tx *sql.Tx, content_id int) {
	query := "DELETE FROM user_content WHERE content_id = ?"
	_, err := tx.ExecContext(ctx, query, content_id)
	if err != nil {
		panic(err)
	}
}