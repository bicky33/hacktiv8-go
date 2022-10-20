package dto

type PhotoCreateRequest struct {
	Title    string `json:"title" validate:"required"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url" validate:"required"`
}

type PhotoResponse struct {
	Code    int32       `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type PhotoCreateResponse struct {
	ID        int32  `json:"id"`
	Title     string `json:"title"`
	Caption   string `json:"caption"`
	PhotoUrl  string `json:"photo_url"`
	UserID    int32  `json:"user_id"`
	CreatedAt string `json:"created_at"`
}

type PhotoGetResponse struct {
	ID        int32             `json:"id"`
	Title     string            `json:"title"`
	Caption   string            `json:"caption"`
	PhotoUrl  string            `json:"photo_url"`
	UserID    int32             `json:"user_id"`
	CreatedAt string            `json:"created_at"`
	UpdatedAt string            `json:"updated_at"`
	User      UserUpdateRequest `json:"User"`
}

type PhotoUpdateResponse struct {
	ID        int32  `json:"id"`
	Title     string `json:"title"`
	Caption   string `json:"caption"`
	PhotoUrl  string `json:"photo_url"`
	UserID    int32  `json:"user_id"`
	UpdatedAt string `json:"updated_at,omitempty"`
}
