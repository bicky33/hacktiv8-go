package dto

type CommentResponse struct {
	Code    int32       `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type CommentCreateRequest struct {
	Message string `json:"message" validate:"required"`
	PhotoId int32  `json:"photo_id" validate:"required"`
}

type CommentCreateResponse struct {
	ID        int32  `json:"id"`
	Message   string `json:"message"`
	PhotoID   int32  `json:"photo_id"`
	UserID    int32  `json:"user_id"`
	CreatedAt string `json:"created_at"`
}

type GetCommentResponse struct {
	ID        int32               `json:"id"`
	UserID    int32               `json:"message"`
	PhotoID   int32               `json:"photo_id"`
	Message   string              `json:"user_id"`
	CreatedAt string              `json:"created_at"`
	UpdatedAt string              `json:"updated_at"`
	User      UserUpdateRequest   `json:"User"`
	Photo     PhotoUpdateResponse `json:"Photo"`
}

type CommentUpdateRequest struct {
	Message string `json:"message" validate:"required"`
}

type CommentUpdateResponse struct {
	ID        int32  `json:"id"`
	Message   string `json:"message"`
	PhotoID   int32  `json:"photo_id"`
	UserID    int32  `json:"user_id"`
	UpdatedAt string `json:"updated_at"`
}
