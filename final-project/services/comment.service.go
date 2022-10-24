package services

import (
	"errors"
	"final-project/dto"
	"final-project/helper"
	"final-project/repositories"
	"fmt"

	"github.com/valyala/fasthttp"
)

type CommentService interface {
	Create(*fasthttp.RequestCtx, dto.CommentCreateRequest, uint32) (*dto.CommentCreateResponse, error)
	GetAll(*fasthttp.RequestCtx) (*[]dto.GetCommentResponse, error)
	Update(*fasthttp.RequestCtx, dto.CommentUpdateRequest, uint32) (*dto.CommentUpdateResponse, error)
	Delete(*fasthttp.RequestCtx, uint32) error
}

type CommentServiceImpl struct {
	DB *repositories.Queries
}

func NewCommentService(DB *repositories.Queries) CommentService {
	return &CommentServiceImpl{DB: DB}
}

func (service *CommentServiceImpl) Create(ctx *fasthttp.RequestCtx, payload dto.CommentCreateRequest, userId uint32) (*dto.CommentCreateResponse, error) {
	data := repositories.InsertCommentParams{
		Message: payload.Message,
		PhotoID: payload.PhotoId,
		UserID:  userId,
	}
	result, err := service.DB.InsertComment(ctx, data)
	if err != nil {
		return nil, err
	}
	createdAt := fmt.Sprintf("%d-%d-%d", result.CreatedAt.Year(), int(result.CreatedAt.Month()), result.CreatedAt.Day())
	responseData := dto.CommentCreateResponse{
		ID:        result.ID,
		Message:   result.Message,
		PhotoID:   result.PhotoID,
		UserID:    result.UserID,
		CreatedAt: createdAt,
	}
	return &responseData, nil
}

func (service *CommentServiceImpl) GetAll(ctx *fasthttp.RequestCtx) (*[]dto.GetCommentResponse, error) {
	result, err := service.DB.GetComment(ctx)
	if err != nil {
		return nil, err
	}
	var responseData []dto.GetCommentResponse
	for _, v := range result {
		user := dto.UserUpdateRequest{
			ID:       v.ID_2,
			Email:    v.Email,
			Username: v.Username,
		}
		photo := dto.PhotoUpdateResponse{
			ID:       v.ID_3,
			Title:    v.Title,
			Caption:  v.Caption,
			PhotoUrl: v.PhotoUrl,
			UserID:   v.UserID_2,
		}
		createdAt := fmt.Sprintf("%d-%d-%d", v.CreatedAt.Year(), int(v.CreatedAt.Month()), v.CreatedAt.Day())
		updatedAt := fmt.Sprintf("%d-%d-%d", v.UpdatedAt.Year(), int(v.UpdatedAt.Month()), v.UpdatedAt.Day())
		comment := dto.GetCommentResponse{
			ID:        v.ID,
			UserID:    v.UserID,
			PhotoID:   v.PhotoID,
			Message:   v.Message,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
			User:      user,
			Photo:     photo,
		}
		responseData = append(responseData, comment)
	}
	return &responseData, nil
}

func (service *CommentServiceImpl) Update(ctx *fasthttp.RequestCtx, payload dto.CommentUpdateRequest, commentId uint32) (*dto.CommentUpdateResponse, error) {
	data := repositories.UpdateCommentParams{
		Message: payload.Message,
		ID:      commentId,
	}
	result, err := service.DB.UpdateComment(ctx, data)
	if err != nil {
		return nil, err
	}
	updatedAt := fmt.Sprintf("%d-%d-%d", result.UpdatedAt.Year(), int(result.UpdatedAt.Month()), result.UpdatedAt.Day())

	responseData := dto.CommentUpdateResponse{
		ID:        result.ID,
		Message:   result.Message,
		PhotoID:   result.PhotoID,
		UserID:    result.UserID,
		UpdatedAt: updatedAt,
	}
	return &responseData, nil
}

func (service *CommentServiceImpl) Delete(ctx *fasthttp.RequestCtx, commentId uint32) error {
	err := service.DB.DeleteComment(ctx, commentId)
	if err != nil {
		return errors.New(helper.MessageDataNotFountError)
	}
	return nil
}
