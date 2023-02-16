package service

import (
	"context"
	"database/sql"
	"task2/helper"
	"task2/model/request"
	"task2/model/response"
	"task2/repository"

	"github.com/go-playground/validator/v10"
)

type userContentService struct {
	db              *sql.DB
	repoContent     repository.ContentRepository
	repoUserContent repository.UserContentRepository
	repoUser        repository.UserRepository
	validator       *validator.Validate
}

func NewUserContentService(db *sql.DB, repoContent repository.ContentRepository, repoUserContent repository.UserContentRepository, repoUser repository.UserRepository, validator_ validator.Validate) *userContentService {
	return &userContentService{
		db:              db,
		repoContent:     repoContent,
		repoUserContent: repoUserContent,
		repoUser:        repoUser,
		validator:       &validator_,
	}
}

func (service *userContentService) PostReward(ctx context.Context, request request.RequestCreateUserContent) (response.ResponseUserContentPoint, error) {
	tx, err := service.db.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)

	// get usercontent contain a user_id and content_id
	userContent := service.repoUserContent.FindByContentIdAndUserId(ctx, tx, request.UserId, request.ContentId)
	// get user
	user := service.repoUser.FindById(ctx, tx, *userContent.GetUserId())

	if *user.GetPoint() < 1 {
		// set point in user_table
		userContent.SetPoint(&request.Point)
		// and then set final point in user table
		totalPointUser := *user.GetPoint()
		totalPointUserContent := *userContent.GetPoint()
		var total = totalPointUser + totalPointUserContent
		user.SetPoint(&total)

		service.repoUser.UpdatePoint(ctx, tx, user)
		service.repoUserContent.UpdatePoint(ctx, tx, userContent)
		return userContent.ToResponseUserContentPoint(), err
	} else {
		// get first and adds first old point value
		point := *userContent.GetPoint()
		userContent.SetPoint(&request.Point)
		pointTotal := point + *userContent.GetPoint()
		// set point in user_table
		userContent.SetPoint(&pointTotal)
		// and then set final point in user table
		totalPointUser := *user.GetPoint()
		totalPointUserContent := *userContent.GetPoint()
		var total = totalPointUser + totalPointUserContent
		user.SetPoint(&total)

		service.repoUser.UpdatePoint(ctx, tx, user)
		service.repoUserContent.UpdatePoint(ctx, tx, userContent)
		return userContent.ToResponseUserContentPoint(), err
	}
}
