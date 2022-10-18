package services

import (
	"errors"
	"final-project/dto"
	"final-project/helper"
	"final-project/repositories"

	"github.com/valyala/fasthttp"
)

type UserService interface {
	Create(*fasthttp.RequestCtx, dto.UserCreateRequest) (repositories.InsertUserRow, error)
	Login(*fasthttp.RequestCtx, dto.UserLogin) (string, error)
}

type UserServiceImpl struct {
	DB *repositories.Queries
}

func NewUserServie(DB *repositories.Queries) UserService {
	return &UserServiceImpl{DB: DB}
}

func (service *UserServiceImpl) Create(ctx *fasthttp.RequestCtx, payload dto.UserCreateRequest) (repositories.InsertUserRow, error) {
	var data = repositories.InsertUserParams{
		Username: payload.Username,
		Password: payload.Password,
		Email:    payload.Email,
		Age:      int32(payload.Age),
	}
	checkEmail, err := service.DB.CountUserEmail(ctx, data.Email)
	if err != nil {
		return repositories.InsertUserRow{}, err
	}
	if checkEmail != 0 {
		return repositories.InsertUserRow{}, errors.New("email must be unique")
	}
	checkUsername, err := service.DB.CountUserUsername(ctx, data.Username)
	if err != nil {
		return repositories.InsertUserRow{}, err
	}
	if checkUsername != 0 {
		return repositories.InsertUserRow{}, errors.New("username must be unique")
	}
	result, err := service.DB.InsertUser(ctx, data)
	return result, err
}

func (service *UserServiceImpl) Login(ctx *fasthttp.RequestCtx, payload dto.UserLogin) (string, error) {
	user, err := service.DB.GetUserEmail(ctx, payload.Email)
	if err != nil {
		return "", err
	}

	if user.Email == "" {
		return "", errors.New("account not found")
	}

	isPasswordMatch := helper.PasswordMatch(user.Password, payload.Password)
	if !isPasswordMatch {
		return "", errors.New("email or password doesnt match")
	}

	return "", nil
}
