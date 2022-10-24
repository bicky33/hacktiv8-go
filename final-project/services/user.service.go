package services

import (
	"errors"
	"final-project/config"
	"final-project/dto"
	"final-project/helper"
	"final-project/repositories"
	"fmt"

	"github.com/valyala/fasthttp"
)

type UserService interface {
	Create(*fasthttp.RequestCtx, dto.UserCreateRequest) (*dto.UserCreateResponse, error)
	Login(*fasthttp.RequestCtx, dto.UserLogin) (string, error)
	Update(*fasthttp.RequestCtx, dto.UserUpdateRequest, uint32) (*dto.UserUpdateResponse, error)
	Delete(*fasthttp.RequestCtx, uint32) error
}

type UserServiceImpl struct {
	DB *repositories.Queries
}

func NewUserServie(DB *repositories.Queries) UserService {
	return &UserServiceImpl{DB: DB}
}

func (service *UserServiceImpl) Create(ctx *fasthttp.RequestCtx, payload dto.UserCreateRequest) (*dto.UserCreateResponse, error) {
	var data = repositories.InsertUserParams{
		Username: payload.Username,
		Password: payload.Password,
		Email:    payload.Email,
		Age:      uint32(payload.Age),
	}
	checkEmail, err := service.DB.CountUserEmail(ctx, data.Email)
	if err != nil {
		return nil, err
	}
	if checkEmail != 0 {
		return nil, errors.New(helper.MessageEmailError)
	}
	checkUsername, err := service.DB.CountUserUsername(ctx, data.Username)
	if err != nil {
		return nil, err
	}
	if checkUsername != 0 {
		return nil, errors.New(helper.MessageUsernameError)
	}
	result, err := service.DB.InsertUser(ctx, data)
	if err != nil {
		return nil, err
	}
	responseData := dto.UserCreateResponse{ID: uint32(result.ID), Age: uint32(result.Age), Username: result.Username, Email: result.Email}
	return &responseData, nil
}

func (service *UserServiceImpl) Login(ctx *fasthttp.RequestCtx, payload dto.UserLogin) (string, error) {

	user, err := service.DB.GetUserByEmail(ctx, payload.Email)
	if err != nil {
		return "", err
	}

	isPasswordMatch := helper.PasswordMatch(user.Password, payload.Password)
	if !isPasswordMatch {
		return "", errors.New(helper.MessagePasswordError)
	}
	config := config.Config()
	data := dto.UserCreateResponse{ID: uint32(user.ID), Username: user.Username, Email: user.Email}
	accessToken, err := helper.GenerateToken(data, config.AccessTokenExpiresIn, config.AccessTokenPrivateKey)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func (service *UserServiceImpl) Update(ctx *fasthttp.RequestCtx, payload dto.UserUpdateRequest, userId uint32) (*dto.UserUpdateResponse, error) {
	data := repositories.UpdateUserParams{ID: userId, Email: payload.Email, Username: payload.Username}
	result, err := service.DB.UpdateUser(ctx, data)
	if err != nil {
		return nil, errors.New(helper.MessageDataNotFountError)
	}
	updatedAt := fmt.Sprintf("%d-%d-%d", result.UpdatedAt.Year(), int(result.UpdatedAt.Month()), result.UpdatedAt.Day())
	responseData := dto.UserUpdateResponse{ID: result.ID, Age: result.Age, Email: result.Email, Username: result.Username, UpdatedAt: updatedAt}
	return &responseData, nil
}

func (service *UserServiceImpl) Delete(ctx *fasthttp.RequestCtx, userId uint32) error {
	err := service.DB.DeleteUser(ctx, userId)
	if err != nil {
		return errors.New(helper.MessageDataNotFountError)
	}
	return nil
}
