package service

import (
	"context"
	"database/sql"
	"sync"
	"task2/helper"
	"task2/model/domain"
	"task2/model/request"
	"task2/model/response"
	"task2/repository"

	"github.com/go-playground/validator/v10"
)

type contentService struct {
	db              *sql.DB
	repoContent     repository.ContentRepository
	repoUserContent repository.UserContentRepository
	repoUser        repository.UserRepository
	validator       *validator.Validate
}

func NewContentService(db *sql.DB, repoContent repository.ContentRepository, repoUserContent repository.UserContentRepository, repoUser repository.UserRepository, validator_ validator.Validate) *contentService {
	return &contentService{
		db:              db,
		repoContent:     repoContent,
		repoUserContent: repoUserContent,
		repoUser:        repoUser,
		validator:       &validator_,
	}
}

func (service *contentService) FindAll(ctx context.Context) ([]response.ResponseUserContentAll, error) {
	tx, err := service.db.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)

	content := service.repoContent.FindAll(ctx, tx)
	responseContent := []response.ResponseUserContentAll{}
	for _, v := range content {
		responseContent = append(responseContent, v.ToResponseUserContentAll())
	}
	return responseContent, nil
}

func (service *contentService) FindById(ctx context.Context, id int) (response.ResponseUserContentByContentId, error) {
	tx, err := service.db.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}
	res := service.repoContent.FindById(ctx, tx, id)
	defer helper.CommitOrRollback(tx)

	return res.ToResponseUserContentByContentId(), nil
}

func (service *contentService) FindContentByUserId(ctx context.Context, user_id, content_id int) (response.ResponseUserContentByContentId, error) {
	tx, err := service.db.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}
	res := service.repoContent.FindContentByContentIdAndUserId(ctx, tx, user_id, content_id)
	defer helper.CommitOrRollback(tx)

	return res.ToResponseUserContentByContentId(), err
}

func (service *contentService) FindByStatusPublish(ctx context.Context) ([]response.ResponseUserContentByContentId, error) {
	tx, err := service.db.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)

	content := service.repoContent.FindByStatusPublish(ctx, tx)
	responseContent := []response.ResponseUserContentByContentId{}
	for _, v := range content {
		responseContent = append(responseContent, v.ToResponseUserContentByContentId())
	}

	return responseContent, err
}

func (service *contentService) FindByStatusDraft(ctx context.Context) ([]response.ResponseUserContentByContentId, error) {
	tx, err := service.db.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)

	content := service.repoContent.FindByStatusDraft(ctx, tx)
	responseContent := []response.ResponseUserContentByContentId{}
	for _, v := range content {
		responseContent = append(responseContent, v.ToResponseUserContentByContentId())
	}

	return responseContent, err
}

func (service *contentService) Create(ctx context.Context, request request.RequestCreateContent, user_id int) (response.ResponseUserContentCreated, error) {
	tx, err := service.db.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)
	dataContent := domain.Content{}
	wg := sync.WaitGroup{}
	wg.Add(4)
	go func() {
		dataContent.SetTitle(&request.Title)
		wg.Done()
	}()

	go func() {
		dataContent.SetBody(&request.Body)
		wg.Done()
	}()

	go func() {
		dataContent.SetImageName(&request.Image_name)
		wg.Done()
	}()

	go func() {
		dataContent.SetStatus(&request.Status)
		wg.Done()
	}()
	wg.Wait()

	// create content and return id content
	content_id := service.repoContent.Create(ctx, tx, dataContent)

	// create user content with user_id from parameter and content_id from create content
	dataUserContent := domain.UserContent{}
	dataUser := domain.User{}
	dataUserContent.SetUserId(&user_id)
	dataUserContent.SetContentId(content_id)
	// set point reward user_content
	var reward int = 10
	dataUserContent.SetPoint(&reward)

	// set point reward user
	totalPointUser := *dataUser.GetPoint()
	dataUser.SetPoint(&reward)
	var total = totalPointUser + *dataUser.GetPoint()
	dataUser.SetPoint(&total)

	service.repoUser.UpdatePoint(ctx, tx, dataUser)
	service.repoUserContent.Create(ctx, tx, dataUserContent)

	return dataContent.ToResponseUserContentCreated(), err
}

func (service *contentService) Update(ctx context.Context, request request.RequestUpdateContent, user_id, content_id int) (response.ResponseUserContentUpdateByContentId, error) {
	tx, err := service.db.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)

	// get usercontent contain a user_id and content_id
	userContent := service.repoUserContent.FindByContentIdAndUserId(ctx, tx, user_id, content_id)
	// if true, get content by content_id
	content := service.repoContent.FindById(ctx, tx, *userContent.GetContentId())
	wg := sync.WaitGroup{}
	wg.Add(4)
	go func() {
		content.SetTitle(&request.Title)
		wg.Done()
	}()

	go func() {
		content.SetBody(&request.Body)
		wg.Done()
	}()

	go func() {
		content.SetImageName(&request.Image_name)
		wg.Done()
	}()

	go func() {
		content.SetStatus(&request.Status)
		wg.Done()
	}()
	wg.Wait()

	// update content
	service.repoContent.Update(ctx, tx, content)
	return content.ToResponseUserContentUpdateByContentId(), err
}

func (service *contentService) Delete(ctx context.Context, user_id, content_id int) error {
	tx, err := service.db.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}
	service.repoContent.Delete(ctx, tx, content_id)
	service.repoUserContent.DeleteByContentId(ctx, tx, user_id, content_id)
	defer helper.CommitOrRollback(tx)
	return nil
}
