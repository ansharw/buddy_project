package repository

import (
	"context"
	"database/sql"
	"task2/model/domain"
)

type contentRepository struct {
}

func NewContentRepository() *contentRepository {
	return &contentRepository{}
}

func (repo *contentRepository) FindAll(ctx context.Context, tx *sql.Tx) []domain.Content {
	query := "SELECT id, title, image_name, body, status, created, modified FROM content ORDER BY id DESC"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	var contents []domain.Content
	for rows.Next() {
		var content domain.Content
		rows.Scan(content.GetId(), content.GetTitle(), content.GetImageName(), content.GetBody(), content.GetStatus(), content.GetCreated(), content.GetModified())
		contents = append(contents, content)
	}
	return contents
}

func (repo *contentRepository) FindById(ctx context.Context, tx *sql.Tx, id int) domain.Content {
	query := "SELECT id, title, image_name, body, status, created, modified FROM content WHERE id = ?"
	content := domain.Content{}
	row := tx.QueryRowContext(ctx, query, id)
	row.Scan(content.GetId(), content.GetTitle(), content.GetImageName(), content.GetBody(), content.GetStatus(), content.GetCreated(), content.GetModified())
	return content
}

func (repo *contentRepository) FindByStatusPublish(ctx context.Context, tx *sql.Tx) []domain.Content {
	query := "SELECT id, title, image_name, body, status, created, modified FROM content WHERE id = 1"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	var contents []domain.Content
	for rows.Next() {
		var content domain.Content
		rows.Scan(content.GetId(), content.GetTitle(), content.GetImageName(), content.GetBody(), content.GetStatus(), content.GetCreated(), content.GetModified())
		contents = append(contents, content)
	}
	return contents
}

func (repo *contentRepository) FindByStatusDraft(ctx context.Context, tx *sql.Tx) []domain.Content {
	query := "SELECT id, title, image_name, body, status, created, modified FROM content WHERE id = 2"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	var contents []domain.Content
	for rows.Next() {
		var content domain.Content
		rows.Scan(content.GetId(), content.GetTitle(), content.GetImageName(), content.GetBody(), content.GetStatus(), content.GetCreated(), content.GetModified())
		contents = append(contents, content)
	}
	return contents
}

func (repo *contentRepository) Create(ctx context.Context, tx *sql.Tx, content domain.Content) *int {
	query := "INSERT INTO content(title, image_name, body, status) VALUES (?, ?, ?, ?)"
	res, err := tx.ExecContext(ctx, query, content.GetTitle(), content.GetImageName(), content.GetBody(), content.GetStatus())
	if err != nil {
		panic(err)
	}
	lastId, _ := res.LastInsertId()
	id := int(lastId)
	content.SetId(&id)
	return content.GetId()
}

// func (repo *contentRepository) Create(ctx context.Context, tx *sql.Tx, content domain.Content) {
// 	query := "INSERT INTO content(IFNULL(NULL, 'No Title'), IFNULL(NULL, 'No Image'), IFNULL(NULL, 'No Description'), status) VALUES (?, ?, ?, ?)"
// 	res, err := tx.ExecContext(ctx, query, content.GetTitle(), content.GetImageName(), content.GetBody(), content.GetStatus())
// 	if err != nil {
// 		panic(err)
// 	}
// 	lastId, _ := res.LastInsertId()
// 	id := int(lastId)
// 	content.SetId(&id)
// }

func (repo *contentRepository) Update(ctx context.Context, tx *sql.Tx, content domain.Content) {
	query := `UPDATE content 
	SET title = 
	CASE WHEN title IS NULL THEN 'No Title'
	ELSE ?
	END AS title,
	image_name = 
	CASE WHEN image_name IS NULL THEN 'No Image'
	ELSE ?
	END AS image_name,
	body = 
	CASE WHEN body IS NULL THEN 'No Description'
	ELSE ?
	END AS body,
	status = 
	CASE WHEN status IS NULL THEN 2
	ELSE ?
	END AS status WHERE id = ?`
	_, err := tx.ExecContext(ctx, query, content.GetTitle(), content.GetImageName(), content.GetBody(), content.GetStatus(), content.GetId())
	if err != nil {
		panic(err)
	}
}

func (repo *contentRepository) Delete(ctx context.Context, tx *sql.Tx, id int) {
	query := "DELETE FROM content WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		panic(err)
	}
}

func (repo *contentRepository) FindContentByContentIdAndUserId(ctx context.Context, tx *sql.Tx, user_id, content_id int) domain.Content {
	query := "SELECT user_content.user_id, content.id, content.title, content.body, content.image_name, content.status FROM user_content JOIN content ON user_content.content_id = content.id WHERE user_id = ? AND content_id = ? ORDER BY content_id DESC"
	userContent := domain.UserContent{}
	content := domain.Content{}
	row := tx.QueryRowContext(ctx, query, user_id, content_id)
	row.Scan(userContent.GetUserId(), content.GetId(), content.GetTitle(), content.GetBody(), content.GetImageName(), content.GetStatus())
	return content
}