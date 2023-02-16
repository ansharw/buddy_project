package repository

import (
	"context"
	"database/sql"
	"task2/model/domain"
)

type userRepository struct {
}

func NewUserRepository() *userRepository {
	return &userRepository{}
}

func (repo *userRepository) FindAll(ctx context.Context, tx *sql.Tx) []domain.User {
	query := "SELECT id, fullname, username, point FROM user ORDER BY id DESC"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	var users []domain.User
	for rows.Next() {
		var user domain.User
		rows.Scan(user.GetId(), user.GetFullname(), user.GetUsername(), user.GetPoint())
		users = append(users, user)
	}
	return users
}

func (repo *userRepository) FindById(ctx context.Context, tx *sql.Tx, id int) domain.User {
	query := "SELECT id, fullname, username, email, point FROM user WHERE id = ?"
	user := domain.User{}
	row := tx.QueryRowContext(ctx, query, id)
	row.Scan(user.GetId(), user.GetFullname(), user.GetUsername(), user.GetEmail(), user.GetPoint())
	return user
}

func (repo *userRepository) FindByUsername(ctx context.Context, tx *sql.Tx, username string) domain.User {
	query := "SELECT password FROM user WHERE username = ?"
	user := domain.User{}
	row := tx.QueryRowContext(ctx, query, username)
	row.Scan(user.GetPassword())
	return user
}

func (repo *userRepository) Login(ctx context.Context, tx *sql.Tx, user domain.User) error {

	query := "SELECT username, password FROM user WHERE username = ? AND password = ?"
	row := tx.QueryRowContext(ctx, query, user.GetUsername(), user.GetPassword())
	err := row.Scan(user.GetUsername(), user.GetPassword())
	if err != nil {
		panic(err)
	}
	return nil
}

func (repo *userRepository) Register(ctx context.Context, tx *sql.Tx, user domain.User) {
	query := "INSERT INTO user(username, email, fullname, password) VALUES(?, ?, ?, ?)"
	res, err := tx.ExecContext(ctx, query, user.GetUsername(), user.GetEmail(), user.GetFullname(), user.GetPassword())
	if err != nil {
		panic(err)
	}
	lastId, _ := res.LastInsertId()
	id := int(lastId)
	user.SetId(&id)
}

func (repo *userRepository) UpdatePoint(ctx context.Context, tx *sql.Tx, user domain.User) {
	query := "UPDATE user SET point = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, user.GetPoint(), user.GetId())
	if err != nil {
		panic(err)
	}
}
