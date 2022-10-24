package dto

type SocialMediaResponse struct {
	Status  uint32      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type SocialMediaCreateRequest struct {
	Name           string `json:"name" validate:"required"`
	SocialMediaUrl string `json:"social_media_url" validate:"required"`
}

type SocialMediaCreateResponse struct {
	ID             uint32 `json:"id"`
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
	UserID         uint32 `json:"user_id"`
	CreatedAt      string `json:"created_at"`
}

type SocialMediaGetResponse struct {
	ID             uint32            `json:"id"`
	Name           string            `json:"name"`
	SocialMediaUrl string            `json:"social_media_url"`
	UserID         uint32            `json:"user_id"`
	CreatedAt      string            `json:"created_at"`
	UpdatedAt      string            `json:"updated_at"`
	User           UserUpdateRequest `json:"User"`
}

type SocialMediaUpdateResponse struct {
	ID             uint32 `json:"id"`
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
	UserID         uint32 `json:"user_id"`
	UpdatedAt      string `json:"updated_at"`
}
