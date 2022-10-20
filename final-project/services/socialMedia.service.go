package services

import (
	"errors"
	"final-project/dto"
	"final-project/repositories"
	"fmt"

	"github.com/valyala/fasthttp"
)

type SocialMedia interface {
	Create(ctx *fasthttp.RequestCtx, payload dto.SocialMediaCreateRequest, userId int32) (*dto.SocialMediaCreateResponse, error)
	GetAll(*fasthttp.RequestCtx) (*[]dto.SocialMediaGetResponse, error)
	Update(*fasthttp.RequestCtx, dto.SocialMediaCreateRequest, int32) (*dto.SocialMediaUpdateResponse, error)
	Delete(*fasthttp.RequestCtx, int32) error
}

type SocialMediaImpl struct {
	DB *repositories.Queries
}

func NewSocialMedia(DB *repositories.Queries) SocialMedia {
	return &SocialMediaImpl{DB: DB}
}

func (service *SocialMediaImpl) Create(ctx *fasthttp.RequestCtx, payload dto.SocialMediaCreateRequest, userId int32) (*dto.SocialMediaCreateResponse, error) {
	data := repositories.InsertSocialMediaParams{
		Name:           payload.Name,
		SocialMediaUrl: payload.SocialMediaUrl,
		UserID:         userId,
	}
	result, err := service.DB.InsertSocialMedia(ctx, data)
	if err != nil {
		return nil, err
	}
	createdAt := fmt.Sprintf("%d-%d-%d", result.CreatedAt.Year(), int(result.CreatedAt.Month()), result.CreatedAt.Day())
	responseData := dto.SocialMediaCreateResponse{
		ID:             result.ID,
		Name:           result.Name,
		SocialMediaUrl: result.SocialMediaUrl,
		UserID:         result.UserID,
		CreatedAt:      createdAt,
	}
	return &responseData, nil
}

func (service *SocialMediaImpl) GetAll(ctx *fasthttp.RequestCtx) (*[]dto.SocialMediaGetResponse, error) {
	result, err := service.DB.GetSocialMedia(ctx)
	if err != nil {
		return nil, err
	}
	var responseData []dto.SocialMediaGetResponse
	for _, v := range result {
		user := dto.UserUpdateRequest{
			ID:              v.ID_2,
			Username:        v.Username,
			ProfileImageUrl: v.ProfileImageUrl,
		}
		createdAt := fmt.Sprintf("%d-%d-%d", v.CreatedAt.Year(), int(v.CreatedAt.Month()), v.CreatedAt.Day())
		updatedAt := fmt.Sprintf("%d-%d-%d", v.UpdatedAt.Year(), int(v.UpdatedAt.Month()), v.UpdatedAt.Day())
		comment := dto.SocialMediaGetResponse{
			ID:             v.ID,
			UserID:         v.UserID,
			SocialMediaUrl: v.SocialMediaUrl,
			Name:           v.Name,
			CreatedAt:      createdAt,
			UpdatedAt:      updatedAt,
			User:           user,
		}
		responseData = append(responseData, comment)
	}
	return &responseData, nil
}

func (service *SocialMediaImpl) Update(ctx *fasthttp.RequestCtx, payload dto.SocialMediaCreateRequest, socialMediaId int32) (*dto.SocialMediaUpdateResponse, error) {
	data := repositories.UpdateSocialMediaParams{
		Name:           payload.Name,
		SocialMediaUrl: payload.SocialMediaUrl,
		ID:             socialMediaId,
	}
	result, err := service.DB.UpdateSocialMedia(ctx, data)
	if err != nil {
		return nil, err
	}
	updatedAt := fmt.Sprintf("%d-%d-%d", result.UpdatedAt.Year(), int(result.UpdatedAt.Month()), result.UpdatedAt.Day())
	responseData := dto.SocialMediaUpdateResponse{
		ID:             result.ID,
		Name:           result.Name,
		SocialMediaUrl: result.SocialMediaUrl,
		UserID:         result.UserID,
		UpdatedAt:      updatedAt,
	}
	return &responseData, nil
}

func (service *SocialMediaImpl) Delete(ctx *fasthttp.RequestCtx, socialMediaId int32) error {
	err := service.DB.DeleteSocialMedia(ctx, socialMediaId)
	if err != nil {
		return errors.New("data not found")
	}
	return nil
}
