package dto

type UserCreateRequest struct {
	Age             int32  `json:"age" validate:"required,numeric,min=8"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=6"`
	Username        string `json:"username" validate:"required"`
	ProfileImageUrl string `json:"profile_image_url" validate:"required"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserCreateResponse struct {
	ID       int32  `json:"id"`
	Age      int32  `json:"age"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UserUpdateRequest struct {
	ID              int32  `json:"id,omitempty"`
	ProfileImageUrl string `json:"profile_image_url,omitempty"`
	Email           string `json:"email" validate:"required,email"`
	Username        string `json:"username" validate:"required"`
}

type UserUpdateResponse struct {
	ID        int32  `json:"id"`
	Age       int32  `json:"age"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	UpdatedAt string `json:"updated_at"`
}

type UserResponse struct {
	Code    int32       `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
