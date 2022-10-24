package dto

type CommentResponse struct {
	Status  uint32      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type CommentCreateRequest struct {
	Message string `json:"message" validate:"required"`
	PhotoId uint32 `json:"photo_id" validate:"required"`
}

type CommentCreateResponse struct {
	ID        uint32 `json:"id"`
	Message   string `json:"message"`
	PhotoID   uint32 `json:"photo_id"`
	UserID    uint32 `json:"user_id"`
	CreatedAt string `json:"created_at"`
}

type GetCommentResponse struct {
	ID        uint32              `json:"id"`
	Message   string              `json:"message"`
	PhotoID   uint32              `json:"photo_id"`
	UserID    uint32              `json:"user_id"`
	CreatedAt string              `json:"created_at"`
	UpdatedAt string              `json:"updated_at"`
	User      UserUpdateRequest   `json:"User"`
	Photo     PhotoUpdateResponse `json:"Photo"`
}

type CommentUpdateRequest struct {
	Message string `json:"message" validate:"required"`
}

type CommentUpdateResponse struct {
	ID        uint32 `json:"id"`
	Message   string `json:"message"`
	PhotoID   uint32 `json:"photo_id"`
	UserID    uint32 `json:"user_id"`
	UpdatedAt string `json:"updated_at"`
}
