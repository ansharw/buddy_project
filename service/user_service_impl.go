package service

import (
	"context"
	"database/sql"
	"task2/helper"
	"task2/model/domain"
	"task2/model/request"
	"task2/model/response"
	"task2/repository"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	db        *sql.DB
	repoUser  repository.UserRepository
	validator *validator.Validate
}

func NewUserService(db *sql.DB, repoUser repository.UserRepository, validator_ validator.Validate) *userService {
	return &userService{
		db:        db,
		repoUser:  repoUser,
		validator: &validator_,
	}
}

func (service *userService) FindAllRegisteredUser(ctx context.Context) ([]response.ResponseRegisteredUser, error) {
	tx, err := service.db.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)

	user := service.repoUser.FindAll(ctx, tx)
	responseUser := []response.ResponseRegisteredUser{}
	for _, v := range user {
		responseUser = append(responseUser, v.ToResponseRegisteredUser())
	}
	return responseUser, nil
}

func (service *userService) FindAllUserPoint(ctx context.Context) ([]response.ResponseRewardPoint, error) {
	tx, err := service.db.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)

	user := service.repoUser.FindAll(ctx, tx)
	responseUser := []response.ResponseRewardPoint{}
	for _, v := range user {
		responseUser = append(responseUser, v.ToResponseRewardPoint())
	}
	return responseUser, nil
}

func (service *userService) FindPointByUserId(ctx context.Context, user_id int) (response.ResponseUserPoint, error) {
	tx, err := service.db.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}
	res := service.repoUser.FindById(ctx, tx, user_id)
	defer helper.CommitOrRollback(tx)

	return res.ToResponseUserPoint(), nil
}

func (service *userService) Login(ctx context.Context, request request.RequestLoginUser) error {
	tx, err := service.db.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)
	dataUser := domain.User{}
	dataUser.SetUsername(&request.Username)
	username := dataUser.GetUsername()
	passwordHashed := service.repoUser.FindByUsername(ctx, tx, *username)
	pass := request.Password
	err = bcrypt.CompareHashAndPassword([]byte(*passwordHashed.GetPassword()), []byte(pass))
	if err != nil {
		panic(err)
	}
	dataUser.SetPassword(passwordHashed.GetPassword())

	err = service.repoUser.Login(ctx, tx, dataUser)
	if err != nil {
		panic(err)
	}
	return nil
}

func (service *userService) Register(ctx context.Context, request request.RequestRegisterUser) error {
	tx, err := service.db.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)
	dataUser := domain.User{}
	dataUser.SetUsername(&request.Username)
	dataUser.SetFullname(&request.Fullname)
	dataUser.SetEmail(&request.Email)

	pass, _ := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	password := string(pass)
	dataUser.SetPassword(&password)

	service.repoUser.Register(ctx, tx, dataUser)
	return nil
}
