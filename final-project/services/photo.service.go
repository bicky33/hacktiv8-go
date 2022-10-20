package services

import (
	"errors"
	"final-project/dto"
	"final-project/repositories"
	"fmt"

	"github.com/valyala/fasthttp"
)

type PhotoService interface {
	Create(*fasthttp.RequestCtx, dto.PhotoCreateRequest, int32) (*dto.PhotoCreateResponse, error)
	GetAll(*fasthttp.RequestCtx) (*[]dto.PhotoGetResponse, error)
	Update(*fasthttp.RequestCtx, dto.PhotoCreateRequest, int32) (*dto.PhotoUpdateResponse, error)
	Delete(*fasthttp.RequestCtx, int32) error
}

type PhotoServiceImpl struct {
	DB *repositories.Queries
}

func NewPhotoService(DB *repositories.Queries) PhotoService {
	return &PhotoServiceImpl{DB: DB}
}

func (service *PhotoServiceImpl) Create(ctx *fasthttp.RequestCtx, payload dto.PhotoCreateRequest, userId int32) (*dto.PhotoCreateResponse, error) {
	data := repositories.InsertPhotoParams{
		Title:    payload.Title,
		Caption:  payload.Caption,
		PhotoUrl: payload.PhotoURL,
		UserID:   userId,
	}
	result, err := service.DB.InsertPhoto(ctx, data)
	if err != nil {
		return nil, err
	}
	createdAt := fmt.Sprintf("%d-%d-%d", result.CreatedAt.Year(), int(result.CreatedAt.Month()), result.CreatedAt.Day())
	responseData := dto.PhotoCreateResponse{
		ID:        result.ID,
		Title:     result.Title,
		Caption:   result.Caption,
		PhotoUrl:  result.PhotoUrl,
		UserID:    result.UserID,
		CreatedAt: createdAt,
	}
	return &responseData, nil
}

func (service *PhotoServiceImpl) GetAll(ctx *fasthttp.RequestCtx) (*[]dto.PhotoGetResponse, error) {
	result, err := service.DB.GetUserPhoto(ctx)
	if err != nil {
		return nil, err
	}
	var responseData []dto.PhotoGetResponse

	for _, v := range result {
		userData := dto.UserUpdateRequest{
			Username: v.Username,
			Email:    v.Email,
		}
		createdAt := fmt.Sprintf("%d-%d-%d", v.CreatedAt.Year(), int(v.CreatedAt.Month()), v.CreatedAt.Day())
		updatedAt := fmt.Sprintf("%d-%d-%d", v.UpdatedAt.Year(), int(v.UpdatedAt.Month()), v.UpdatedAt.Day())
		data := dto.PhotoGetResponse{
			ID:        v.ID,
			Title:     v.Title,
			Caption:   v.Caption,
			PhotoUrl:  v.PhotoUrl,
			UserID:    v.UserID,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
			User:      userData,
		}
		responseData = append(responseData, data)
	}
	return &responseData, nil
}

func (service *PhotoServiceImpl) Update(ctx *fasthttp.RequestCtx, payload dto.PhotoCreateRequest, photoId int32) (*dto.PhotoUpdateResponse, error) {
	data := repositories.UpdatePhotoParams{
		Title:    payload.Title,
		Caption:  payload.Caption,
		PhotoUrl: payload.PhotoURL,
		ID:       photoId,
	}
	result, err := service.DB.UpdatePhoto(ctx, data)
	if err != nil {
		return nil, errors.New("data not found")
	}
	updatedAt := fmt.Sprintf("%d-%d-%d", result.UpdatedAt.Year(), int(result.UpdatedAt.Month()), result.UpdatedAt.Day())
	responseData := dto.PhotoUpdateResponse{
		ID:        result.ID,
		Title:     result.Title,
		Caption:   result.Caption,
		UserID:    result.UserID,
		PhotoUrl:  result.PhotoUrl,
		UpdatedAt: updatedAt,
	}
	return &responseData, nil
}

func (service *PhotoServiceImpl) Delete(ctx *fasthttp.RequestCtx, photoId int32) error {
	err := service.DB.DeletePhoto(ctx, photoId)
	if err != nil {
		return errors.New("data not found")
	}
	return nil
}
